package services

import (
	"errors"
	"fmt"

	"social_network/internal/domain/models"
	"social_network/internal/domain/ports/repository"
)

type FollowServiceImpl struct {
	followRepo repository.FollowRepository
	userRepo   repository.UserRepository
}

func NewFollowService(followRepo repository.FollowRepository, userRepo repository.UserRepository) *FollowServiceImpl {
	return &FollowServiceImpl{followRepo: followRepo, userRepo: userRepo}
}

func (s *FollowServiceImpl) CreateFollow(follow *models.Follow) error {
	if follow.FollowerID == 0 || follow.FollowingID == 0 {
		return errors.New("follower and following IDs must be provided")
	}

	if follow.FollowerID == follow.FollowingID {
		return errors.New("you cannot follow yourself")
	}

	user, err := s.userRepo.GetUserByID(follow.FollowingID)
	if err != nil {
		return fmt.Errorf("user not found: %w", err)
	}

	if user.IsPublic {
		follow.Status = "accepted"
	} else {
		follow.Status = "pending"
	}

	return s.followRepo.CreateFollow(follow)
}
