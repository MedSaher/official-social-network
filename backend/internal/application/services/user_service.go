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

func NewUserService(userRepo repository.UserRepository, followRepo repository.FollowRepository,postRepo repository.PostRepository) *UserServiceImpl {
	return &UserServiceImpl{userRepo: userRepo, followRepo: followRepo,postRepo: postRepo}
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
	fmt.Println("ffff", userID)
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
	fmt.Println("cccccccccc",profileOwnerID)
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

	//get posts user
	posts, err := s.postRepo.GetPostsByUserID(profileOwnerID)
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

func (s *UserServiceImpl) ChangePrivacyStatus(userID int, privacyStatus string) error {
	valid := map[string]bool{
		"public": true, "private": true, "almost_private": true,
	}
	if !valid[privacyStatus] {
		return errors.New("invalid privacy status")
	}
	return s.userRepo.UpdatePrivacyStatus(userID, privacyStatus)
}

func (s *UserServiceImpl) SearchUsers(query string) ([]models.UserProfileDTO, error) {
	return s.userRepo.SearchUsers(query)
}

func (s *UserServiceImpl) GetUserProfileByUsername(username string) (*models.UserProfileDTO, error) {
	return s.userRepo.GetUserProfileByUsername(username)
}

func (s *UserServiceImpl) GetAnotherProfile(viewerID int, profileOwnerUsername string) (*models.AnotherProfileResponse, error) {
	user, err := s.userRepo.GetUserByUsername(profileOwnerUsername)
	if err != nil {
		return nil, err
	}

	followers, err := s.followRepo.GetFollowers(user.Id)
	if err != nil {
		return nil, err
	}
	following, err := s.followRepo.GetFollowing(user.Id)
	if err != nil {
		return nil, err
	}

	isSelf := viewerID == user.Id

	isFollowing := false
	status, err := s.followRepo.GetStatusFollow(viewerID, user.Id)
	if err == nil && status == "accepted" {
		isFollowing = true
	}

	canSeePosts := user.PrivacyStatus == "public" || isSelf || isFollowing

	var posts []models.Post
	if canSeePosts {
		posts, _ = s.postRepo.GetPostsByUserID(user.Id)
	}

	resp := &models.AnotherProfileResponse{
		Id:             user.Id,
		UserName:       user.UserName,
		FirstName:      "",
		LastName:       "",
		AvatarUrl:      user.AvatarPath,
		AboutMe:        nil,
		PrivacyStatus:  user.PrivacyStatus,
		FollowersCount: len(followers),
		FollowingCount: len(following),
		Posts:          posts,
		IsSelf:         isSelf,
		IsFollowing:    isFollowing,
	}

	if canSeePosts {
		resp.FirstName = user.FirstName
		resp.LastName = user.LastName
		resp.AboutMe = user.AboutMe
	}

	return resp, nil
}
