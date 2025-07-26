package service

import (
	"context"
	"social_network/internal/domain/models"
)

type PostService interface {
	CreatePost(ctx context.Context, userID int, groupID *int, content, privacy, imagePath string) (models.Post, error)
	GetAllPosts(ctx context.Context) ([]models.Post, error)
	GetCommentsByPostID(ctx context.Context, postID int) ([]models.Comment, error)
    CreateComment(ctx context.Context, c *models.Comment) error
}
