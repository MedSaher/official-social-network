package service

import "social_network/internal/domain/models"

type FollowService interface {
	CreateFollow(follow *models.Follow) error
}
