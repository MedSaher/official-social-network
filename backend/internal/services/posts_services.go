package services

import (
	"errors"
	"fmt"
	"time"

	"social_network/internal/models"
	"social_network/internal/repositories"
)

// Create an interface for the posts services:
type PostsServiceLayer interface {
	CreatePost(post *models.PostUser, token string) error
	GetAllPostsService(offset, limit int) ([]*models.PostUser, error) 
	GetAllCategoriesService() ([]*models.Categories, error)
}

// Create a structure to implement the functionalities
// within our interface contract:
type PostsService struct {
	PostRepo    repositories.PostsRepositoryLayer
	SessionRepo repositories.SessionsRepositoryLayer
}

// Instantiate a new postService instance:
func NewPostService(postRepo *repositories.PostsRepository, sessRepo *repositories.SessionsRepository) *PostsService {
	return &PostsService{
		PostRepo:    postRepo,
		SessionRepo: sessRepo,
	}
}

// Create a new post server:
func (postSer *PostsService) CreatePost(post *models.PostUser, token string) error {
	var err error
	if post.Title == "" || post.Content == "" {
		return errors.New("missing content or title")
	}
	post.CreatedAt = time.Now().UTC().Format("2006-01-02 15:04:05")
	post.UserId, err = postSer.SessionRepo.GetSessionByToken(token)
	if err != nil || post.UserId == 0{
		if err != nil {
				fmt.Printf("error retreiving user id ->----------->: %v", err)
			return err
		}
		fmt.Printf("error retreiving user id.")
		return  errors.New("error retreiving user id.")
	}
	return postSer.PostRepo.CreatePost(post)
}

// Get all posts service:
func (postSer *PostsService) GetAllPostsService(offset, limit int) ([]*models.PostUser, error) {
	posts, err := postSer.PostRepo.GetAllPostsRepository(offset, limit)
	if err != nil {
		return nil, err
	}

	for _, post := range posts {
		parsedTime, err := time.Parse("2006-01-02 15:04:05", post.CreatedAt)
		if err == nil {
			post.CreatedAt = parsedTime.Format(time.RFC3339)
		}
	}
	return posts, nil
}


// Get all categories service:
func (postSer *PostsService) GetAllCategoriesService() ([]*models.Categories, error) {
	categ, err := postSer.PostRepo.GetCategories()
	if err != nil {
		return nil, err
	}
	return categ, nil
}
