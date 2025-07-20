package service

import "social_network/internal/domain/models"

type UserService interface {
	Register(user *models.User) error
	GetByUsername(username string) (*models.User, error)
}
