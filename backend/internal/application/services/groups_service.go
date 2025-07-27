package services

import (
	"context"
	"fmt"

	"social_network/internal/domain/models"
	"social_network/internal/domain/ports/repository"
	"social_network/internal/domain/ports/service"
)

type groupService struct {
	repo repository.GroupRepository
}

func NewGroupService(r repository.GroupRepository) service.GroupService {
	return &groupService{repo: r}
}

func (s *groupService) CreateGroup(ctx context.Context, g *models.Group) error {
	return s.repo.CreateGroup(ctx, g)
}

func (s *groupService) GetGroupsForUser(ctx context.Context, userID int) ([]models.GroupWithUserFlags, error) {
	return s.repo.GetAllGroupsForUser(ctx, userID)
}

func (s *groupService) RequestToJoinGroup(ctx context.Context, groupID int, userID int) error {
	// Check for existing membership
	exists, err := s.repo.IsAlreadyMember(ctx, groupID, userID)
	if err != nil {
		return fmt.Errorf("failed to check existing membership: %v", err)
	}
	if exists {
		return fmt.Errorf("already requested or member")
	}

	// Insert join request
	return s.repo.CreateJoinRequest(ctx, groupID, userID)
}
