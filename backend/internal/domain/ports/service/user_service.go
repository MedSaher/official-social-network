package service

import "social_network/internal/domain/models"

type UserService interface {
	Register(user *models.User) error
	Authenticate(email, password string) (*models.User, error)
	GetProfile(id int) (*models.User, error)
	GetFullProfile(userID int) (*models.FullProfileResponse, error)
	GetFullProfileData(viewerID, profileOwnerID int) (*models.FullProfileResponse, error) // ✅ NEW
}
