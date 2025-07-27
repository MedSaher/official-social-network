package repository

import (
	"context"
	"social_network/internal/domain/models"
)

type GroupRepository interface {
	CreateGroup(ctx context.Context, g *models.Group) error
	GetAllGroupsForUser(ctx context.Context, userID int) ([]models.GroupWithUserFlags, error)
	IsAlreadyMember(ctx context.Context, groupID int, userID int) (bool, error)
	CreateJoinRequest(ctx context.Context, groupID int, userID int) error
	IsCreator(ctx context.Context, groupID, userID int) (bool, error)
	GetPendingRequests(ctx context.Context, groupID int) ([]models.GroupJoinRequest, error)
}
