package repository

import (
	"context"
	"social_network/internal/domain/models"
)

type GroupRepository interface {
	CreateGroup(ctx context.Context, g *models.Group) error
}
