package repository

import "social_network/internal/domain/models"

type UserRepository interface {
	RegisterNewUser(user *models.User) error
	GetUserByUsername(username string) (*models.User, error)
}