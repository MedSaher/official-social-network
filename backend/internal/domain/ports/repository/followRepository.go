package repository

type FollowRepository interface {
	// Create(follow *Follow) error
	Accept(followerID, followingID int) error
	Decline(followerID, followingID int) error
	Delete(followerID, followingID int) error
	GetStatus(followerID, followingID int) (string, error)
	// GetFollowers(userID int) ([]Follow, error)
	// GetFollowing(userID int) ([]Follow, error)
}