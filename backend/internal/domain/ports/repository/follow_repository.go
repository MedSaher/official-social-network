package repository

import "social_network/internal/domain/models"

type FollowRepository interface {
	CreateFollow(follow *models.Follow) error
	AcceptFollow(followerID, followingID int) error
	DeclineFollow(followerID, followingID int) error
	DeleteFollow(followerID, followingID int) error
	GetStatusFollow(followerID, followingID int) (string, error)
	GetFollowers(userID int) ([]models.FollowerInfo, error)
	GetFollowing(userID int) ([]models.FollowerInfo, error)
	IsFollowing(followerID, followingID int) (bool, error)
}
