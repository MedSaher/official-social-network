package services

import (
	"errors"
	"fmt"
	"time"

	"social_network/internal/models"
	"social_network/internal/repositories"
)

// PostsServiceLayer defines the interface for post-related business logic
type PostsServiceLayer interface {
	CreatePost(post *models.PostUser, token string) error
	GetAllPostsService(offset, limit int) ([]*models.PostUser, error)
	GetAllCategoriesService() ([]*models.Categories, error)
}

// PostsService implements PostsServiceLayer
type PostsService struct {
	PostRepo    repositories.PostsRepositoryLayer
	SessionRepo repositories.SessionsRepositoryLayer
}

// NewPostService returns a new PostsService instance
func NewPostService(postRepo *repositories.PostsRepository, sessRepo *repositories.SessionsRepository) *PostsService {
	return &PostsService{
		PostRepo:    postRepo,
		SessionRepo: sessRepo,
	}
}

// CreatePost validates input, attaches user ID, sets time, and creates the post
func (postSer *PostsService) CreatePost(post *models.PostUser, token string) error {
	if post.Title == "" || post.Content == "" {
		return errors.New("missing title or content")
	}

	post.CreatedAt = time.Now().UTC().Format("2006-01-02 15:04:05")

	userID, err := postSer.SessionRepo.GetSessionByToken(token)
	if err != nil || userID == 0 {
		if err != nil {
			fmt.Printf("error retrieving user id from session: %v\n", err)
			return err
		}
		fmt.Println("user id is 0 — invalid session?")
		return errors.New("unable to retrieve user from session")
	}

	post.UserId = userID

	return postSer.PostRepo.CreatePost(post)
}

// GetAllPostsService fetches posts and formats created_at time to RFC3339
func (postSer *PostsService) GetAllPostsService(offset, limit int) ([]*models.PostUser, error) {
	posts, err := postSer.PostRepo.GetAllPostsRepository(offset, limit)
	if err != nil {
		return nil, err
	}

	for _, post := range posts {
		if parsedTime, err := time.Parse("2006-01-02 15:04:05", post.CreatedAt); err == nil {
			post.CreatedAt = parsedTime.UTC().Format(time.RFC3339)
		}
	}

	return posts, nil
}

// GetAllCategoriesService returns all categories
func (postSer *PostsService) GetAllCategoriesService() ([]*models.Categories, error) {
	return postSer.PostRepo.GetCategories()
}
