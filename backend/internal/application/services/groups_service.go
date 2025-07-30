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

func (s *groupService) IsCreator(ctx context.Context, groupID, userID int) (bool, error){
	return s.repo.IsCreator(ctx, groupID, userID)
}

func (s *groupService) 	GetPendingRequests(ctx context.Context, groupID int) ([]models.GroupJoinRequest, error){
	return s.repo.GetPendingRequests(ctx, groupID)
}

// internal/application/services/group_service.go

func (s *groupService) RespondToJoinRequest(ctx context.Context, requestID int, actorID int, accept bool) error {
	// Get the join request
	request, err := s.repo.GetGroupMemberByID(ctx, requestID)
	if err != nil {
		return fmt.Errorf("join request not found")
	}

	// Check if actor is the creator of the group
	isCreator, err := s.repo.IsUserGroupCreator(ctx, actorID, request.GroupID)
	if err != nil {
		return fmt.Errorf("unable to check group creator: %v", err)
	}
	if !isCreator {
		return fmt.Errorf("only group creator can respond to requests")
	}

	// Update request status
	newStatus := "declined"
	if accept {
		newStatus = "accepted"
	}

	err = s.repo.UpdateGroupMemberStatus(ctx, requestID, newStatus)
	if err != nil {
		return fmt.Errorf("failed to update request: %v", err)
	}

	return nil
}

func (s *groupService) GetUserRole(ctx context.Context, groupID, userID int) (string, error) {
	return s.repo.GetUserRole(ctx, groupID, userID)
}

func (s *groupService) GetGroupPosts(ctx context.Context, groupID int) ([]models.GroupPost, error) {
	return s.repo.GetGroupPosts(ctx, groupID)
}

func (s *groupService) GetGroupEvents(ctx context.Context, groupID int) ([]models.GroupEvent, error) {
	return s.repo.GetGroupEvents(ctx, groupID)
}