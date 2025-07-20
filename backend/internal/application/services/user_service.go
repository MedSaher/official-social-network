package services

import (
	"errors"

	"social_network/internal/domain/models"
	"social_network/internal/domain/ports/repository"
	"social_network/internal/domain/ports/service"
)

type userServiceImpl struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) service.UserService {
	return &userServiceImpl{repo: repo}
}

func (s *userServiceImpl) Register(user *models.User) error {
	if user.Email == "" || user.Password == "" {
		return errors.New("missing credentials")
	}
	return s.repo.RegisterNewUser(user)
}

func (s *userServiceImpl) GetByUsername(username string) (*models.User, error) {
	return s.repo.GetUserByUsername(username)
}
