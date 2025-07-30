package service

import (
	"context"
	"social_network/internal/domain/models"
)

type GroupService interface {
	CreateGroup(ctx context.Context, g *models.Group) error
	GetGroupsForUser(ctx context.Context, userID int) ([]models.GroupWithUserFlags, error)
	RequestToJoinGroup(ctx context.Context, groupID int, userID int) error
	IsCreator(ctx context.Context, groupID, userID int) (bool, error)
	GetPendingRequests(ctx context.Context, groupID int) ([]models.GroupJoinRequest, error)
	RespondToJoinRequest(ctx context.Context, requestID int, actorID int, accept bool) error
	GetUserRole(ctx context.Context, groupID, userID int) (string, error)
	GetGroupEvents(ctx context.Context, groupID int) ([]models.GroupEvent, error)
	GetGroupPosts(ctx context.Context, groupID int) ([]models.GroupPost, error)
	GetGroupInfo(ctx context.Context, groupID int) (*models.GroupInfo, error)
	CreateGroupEvent(ctx context.Context, event *models.GroupEvent) error
}
