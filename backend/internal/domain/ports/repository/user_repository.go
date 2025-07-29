package repository

import (
	"social_network/internal/domain/models"
)

type UserRepository interface {
	RegisterNewUser(user *models.User) error
	GetUserByEmail(email string) (*models.User, error)
	GetUserByID(id int) (*models.User, error)
	UpdatePrivacyStatus(userID int, privacyStatus string) error
	SearchUsers(query string) ([]models.UserProfileDTO, error)
	GetUserProfileByUsername(username string) (*models.UserProfileDTO, error)
	GetUserByUsername(username string) (*models.User, error) 
}
