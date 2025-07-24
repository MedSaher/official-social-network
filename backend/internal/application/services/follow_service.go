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

	if user.PrivacyStatus == "public" {
		follow.Status = "accepted"
	} else {
		follow.Status = "pending"
	}

	return s.followRepo.CreateFollow(follow)
}

func (s *FollowServiceImpl) AcceptFollow(followerID, followingID, currentUserID int) error {
	if followerID == 0 || followingID == 0 {
		return errors.New("follower and following IDs must be provided")
	}

	//check if the follower and following IDs are not the same
	if followerID == followingID {
		return errors.New("you cannot accept a follow request from yourself")
	}

	// check if the current user is the one who sent the follow request
	if currentUserID != followingID {
		return errors.New("you are not authorized to accept this follow request")
	}

	return s.followRepo.AcceptFollow(followerID, followingID)
}

func (s *FollowServiceImpl) DeclineFollow(followerID, followingID, currentUserID int) error {
	if followerID == 0 || followingID == 0 {
		return errors.New("follower and following IDs must be provided")
	}

	// check if the follower and following IDs are not the same
	if followerID == followingID {
		return errors.New("you cannot decline a follow request from yourself")
	}
	// check if the current user is the one who sent the follow request
	if currentUserID != followingID {
		return errors.New("you are not authorized to decline this follow request")
	}

	return s.followRepo.DeclineFollow(followerID, followingID)
}

func (s *FollowServiceImpl) DeleteFollow(followerID, followingID, currentUserID int) error {
	if followerID == 0 || followingID == 0 {
		return errors.New("follower and following IDs must be provided")
	}

	// check if the follower and following IDs are not the same
	if followerID == followingID {
		return errors.New("you cannot delete a follow relationship with yourself")
	}

	// check if the current user is the one who sent the follow request
	if currentUserID != followerID && currentUserID != followingID {
		return errors.New("you are not authorized to delete this follow relationship")
	}

	return s.followRepo.DeleteFollow(followerID, followingID)
}

func (s *FollowServiceImpl) GetStatusFollow(followerID, followingID int) (string, error) {
	if followerID == 0 || followingID == 0 {
		return "", errors.New("follower and following IDs must be provided")
	}

	return s.followRepo.GetStatusFollow(followerID, followingID)
}

func (s *FollowServiceImpl) GetFollowers(userID int) ([]models.FollowerInfo, error) {
	if userID == 0 {
		return nil, errors.New("user ID must be provided")
	}

	return s.followRepo.GetFollowers(userID)
}

func (s *FollowServiceImpl) GetFollowing(userID int) ([]models.FollowerInfo, error) {
	if userID == 0 {
		return nil, errors.New("user ID must be provided")
	}

	return s.followRepo.GetFollowing(userID)
}