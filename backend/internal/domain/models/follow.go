package models

type Follow struct {
	ID           int
	FollowerID   int
	FollowingID  int
	Status       string
	CreatedAt    string
	UpdatedAt    string
}