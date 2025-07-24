package service

import "social_network/internal/domain/models"

type FollowService interface {
	CreateFollow(follow *models.Follow) error
	AcceptFollow(followerID, followingID, currentUserID int) error
	DeclineFollow(followerID, followingID, currentUserID int) error
	DeleteFollow(followerID, followingID, currentUserID int) error
}
