package repository

import (
	"context"
	"social_network/internal/domain/models"
)

type PostRepository interface {
	CreatePost(ctx context.Context, userID int, groupID *int, content, privacy, imagePath string) (models.Post, error)
}
