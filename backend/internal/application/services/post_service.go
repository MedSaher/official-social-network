package services

import (
    "context"
    "social_network/internal/domain/ports/service"
    "social_network/internal/domain/ports/repository"
    "social_network/internal/domain/models"
)

type postService struct {
    postRepo repository.PostRepository
}

func NewPostService(postRepo repository.PostRepository) service.PostService {
    return &postService{
        postRepo: postRepo,
    }
}

func (s *postService) CreatePost(ctx context.Context, userID int, groupID *int, content, privacy, imagePath string) (models.Post, error) {
    // Here you can add extra business logic, validation, or preprocessing if needed
    return s.postRepo.CreatePost(ctx, userID, groupID, content, privacy, imagePath)
}

func (s *postService) GetAllPosts(ctx context.Context) ([]models.Post, error) {
    return s.postRepo.GetAllPosts(ctx)
}

func (s *postService) GetCommentsByPostID(ctx context.Context, postID int) ([]models.Comment, error) {
    return s.postRepo.GetCommentsByPostID(ctx, postID)
}

func (s *postService) CreateComment(ctx context.Context, c *models.Comment) error {
    return s.postRepo.CreateComment(ctx, c)
}
