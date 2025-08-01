package service

import "social_network/internal/domain/models"

type FollowService interface {
	CreateFollow(follow *models.Follow) error
	AcceptFollow(followerID, followingID, currentUserID int) error
	DeclineFollow(followerID, followingID, currentUserID int) error
	DeleteFollow(followerID, followingID, currentUserID int) error
	GetStatusFollow(followerID, followingID int) (string, error)
	GetFollowers(userID int) ([]models.FollowerInfo, error)
	GetFollowing(userID int) ([]models.FollowerInfo, error)
}
