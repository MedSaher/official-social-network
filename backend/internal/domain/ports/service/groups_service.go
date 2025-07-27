package service

import (
	"context"
	"social_network/internal/domain/models"
)

type GroupService interface {
	CreateGroup(ctx context.Context, g *models.Group) error
	GetGroupsForUser(ctx context.Context, userID int) ([]models.GroupWithUserFlags, error)
	RequestToJoinGroup(ctx context.Context, groupID int, userID int) error
}
