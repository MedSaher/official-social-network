package services

import (
	"errors"
	"fmt"

	"social_network/internal/application/utils"
	"social_network/internal/domain/models"
	"social_network/internal/domain/ports/repository"
)

type UserServiceImpl struct {
	userRepo repository.UserRepository
}

func NewUserService(userRepo repository.UserRepository) *UserServiceImpl {
	return &UserServiceImpl{userRepo: userRepo}
}

func (s *UserServiceImpl) Register(user *models.User) error {
	fmt.Println("User to register:", user)

	hashedPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		return err
	}

	user.NickName = string(user.LastName[0]) + user.FirstName
	user.Password = hashedPassword

	return s.userRepo.RegisterNewUser(user)
}

func (s *UserServiceImpl) Authenticate(email, password string) (*models.User, error) {
	if email == "" || password == "" {
		return nil, errors.New("email and password required")
	}

	user, err := s.userRepo.GetUserByEmail(email)
	if err != nil || !utils.CheckPasswordHash(password, user.Password) {
		return nil, errors.New("invalid credentials")
	}

	return user, nil
}

func (s *UserServiceImpl) GetProfile(id int) (*models.User, error) {
	return s.userRepo.GetUserByID(id)
}
