package service

import (
	"context"
	"social_network/internal/domain/models"
)

type GroupService interface {
	CreateGroup(ctx context.Context, g *models.Group) error
	GetAllGroups(ctx context.Context) ([]models.Group, error)
}
