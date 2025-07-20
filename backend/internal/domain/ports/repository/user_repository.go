package repository

import (
	"social_network/internal/domain/models"
)

type UserRepository interface {
	RegisterNewUser(user *models.User) error
	GetUserByEmail(email string) (*models.User, error)
	GetUserByID(id int) (*models.User, error)
}
