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
	GetGroupMemberByID(ctx context.Context, requestID int) (*models.GroupMember, error)
	IsUserGroupCreator(ctx context.Context, userID int, groupID int) (bool, error)
	UpdateGroupMemberStatus(ctx context.Context, requestID int, newStatus string) error
	GetUserRole(ctx context.Context, groupID, userID int) (string, error)
	GetGroupPosts(ctx context.Context, groupID int) ([]models.GroupPost, error)
	GetGroupEvents(ctx context.Context, groupID int) ([]models.GroupEvent, error)
	GetGroupInfo(ctx context.Context, groupID int) (*models.GroupInfo, error)
}
