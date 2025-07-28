package services

import (
	"errors"
	"fmt"

	"social_network/internal/application/utils"
	"social_network/internal/domain/models"
	"social_network/internal/domain/ports/repository"
)

type UserServiceImpl struct {
	userRepo   repository.UserRepository
	followRepo repository.FollowRepository
	postRepo   repository.PostRepository
}

func NewUserService(userRepo repository.UserRepository, followRepo repository.FollowRepository) *UserServiceImpl {
	return &UserServiceImpl{userRepo: userRepo, followRepo: followRepo}
}

func (s *UserServiceImpl) Register(user *models.User) error {
	fmt.Println("User to register:", user)

	hashedPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		return err
	}

	user.UserName = string(user.LastName[0]) + user.FirstName
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

func (s *UserServiceImpl) GetFullProfile(userID int) (*models.FullProfileResponse, error) {
	fmt.Println("ffff",userID)
	user, err := s.userRepo.GetUserByID(userID)
	if err != nil {
		return nil, err
	}

	followers, err := s.followRepo.GetFollowers(userID)
	if err != nil {
		return nil, err
	}

	following, err := s.followRepo.GetFollowing(userID)
	if err != nil {
		return nil, err
	}

	//get posts user
	posts, err := s.postRepo.GetPostsByUserID(userID)
	if err != nil {
		return nil, err
	}

	userDTO := models.UserProfileDTOFromUser(user)

	return &models.FullProfileResponse{
		User:           userDTO,
		FollowersCount: len(followers),
		FollowingCount: len(following),
		Posts:          posts,
	}, nil
}

// ✅ NEW: GetFullProfileData with viewerID (could be used later for isFollowing logic)
func (s *UserServiceImpl) GetFullProfileData(viewerID, profileOwnerID int) (*models.FullProfileResponse, error) {
	user, err := s.userRepo.GetUserByID(profileOwnerID)
	if err != nil {
		return nil, err
	}

	followers, err := s.followRepo.GetFollowers(profileOwnerID)
	if err != nil {
		return nil, err
	}

	following, err := s.followRepo.GetFollowing(profileOwnerID)
	if err != nil {
		return nil, err
	}

	userDTO := models.UserProfileDTOFromUser(user)

	return &models.FullProfileResponse{
		User:           userDTO,
		FollowersCount: len(followers),
		FollowingCount: len(following),
	}, nil
}

func (s *UserServiceImpl) SearchUsers(query string) ([]models.UserProfileDTO, error) {
	return s.userRepo.SearchUsers(query)
}

func (s *UserServiceImpl) GetUserProfileByUsername(username string) (*models.UserProfileDTO, error) {
	return s.userRepo.GetUserProfileByUsername(username)
}
