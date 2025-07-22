package models

type Follow struct {
	ID           int    `json:"id"`
	FollowerID   int    `json:"follower_id"`
	FollowingID  int    `json:"following_id"`
	Status       string `json:"status"`
	CreatedAt    string `json:"created_at"`
	UpdatedAt    string `json:"updated_at"`
}