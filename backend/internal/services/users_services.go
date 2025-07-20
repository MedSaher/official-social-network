package services

import (
	"errors"
	"fmt"

	"social_network/internal/models"
	"social_network/internal/repositories"
	"social_network/internal/services/utils"
)

// Create an interface to describe the functionalities of the user services:
type UsersServicesLayer interface {
	UserRegestration(user *models.User) error
	AuthenticateUser(email, password string) (*models.User, error)
	// GetUsersService(offset, limit int) ([]*models.ChatUser, error)
	GetUserProfile(userId int) (*models.User, error)
}

// Create structure to implement the services innterfase:
type UsersServices struct {
	userRepository repositories.UsersRepositoryLayer
}

// Instantiate the user_service structure:
func NewUsersServices(userRepo repositories.UsersRepositoryLayer) *UsersServices {
	return &UsersServices{userRepository: userRepo}
}

// Register q new user service:
func (userServ *UsersServices) UserRegestration(user *models.User) error {
	// if user.FirstName == "" || user.LastName == "" || user.Email == "" || !utils.IsValidGender(user.Gender) || user.PasswordHash == "" {
	// 	return errors.New("invalid credentials")
	// }
	fmt.Println("The new registred user: ----> ", user)
	hashedPassword, err := utils.HashPassword(user.PasswordHash)
	if err != nil {
		return err
	}
	user.Nickname = string(user.LastName[0]) + user.FirstName
	user.PasswordHash = hashedPassword
	return userServ.userRepository.RegisterNewUser(user)
}

// AuthenticateUser verifies user credentials and returns the user if valid
func (userServ *UsersServices) AuthenticateUser(email, password string) (*models.User, error) {
	// Input validation
	if email == "" {
		return nil, errors.New("email is required")
	}
	if password == "" {
		return nil, errors.New("password is required")
	}

	// Get user by email
	user, err := userServ.userRepository.GetUserByEmail(email)
	if err != nil {
		// Log the error but don't expose details to client
		return nil, errors.New("invalid email or password")
	}

	// Check if password matches
	if !utils.CheckPasswordHash(password, user.PasswordHash) {
		return nil, errors.New("invalid email or password")
	}

	return user, nil
}

// // Get all users to fill the chat menu (removed offset and limit):
// func (userServ *UsersServices) GetUsersService(offset, limit int) ([]*models.ChatUser, error) {
// 	return userServ.userRepository.GetUsersRepo(offset, limit)
// }

// extract the user from dataabase:
func (userServ *UsersServices) GetUserProfile(userId int) (*models.User, error) {
	return userServ.userRepository.GetUserByID(userId)
}
