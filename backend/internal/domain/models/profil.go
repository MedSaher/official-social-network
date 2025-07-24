package models

type FullProfileResponse struct {
	User           UserProfileDTO `json:"user"`
	FollowersCount int            `json:"followers_count"`
	FollowingCount int            `json:"following_count"`
	Followers      []FollowerInfo `json:"followers"`
	Following      []FollowerInfo `json:"following"`
}


