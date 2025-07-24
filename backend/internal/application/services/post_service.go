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
