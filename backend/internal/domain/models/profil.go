package models

type FullProfileResponse struct {
	User           UserProfileDTO `json:"user"`
	FollowersCount int            `json:"followers_count"`
	FollowingCount int            `json:"following_count"`
	Posts          []Post 
}

type PublicProfileResponse struct {
    Id            int     `json:"id"`
    UserName      string  `json:"username"`
    FirstName     string  `json:"firstName,omitempty"`   // hide if private
    LastName      string  `json:"lastName,omitempty"`
    AvatarUrl     *string `json:"avatarUrl,omitempty"`
    AboutMe       *string `json:"aboutMe,omitempty"`
    PrivacyStatus string  `json:"privacyStatus"`
    FollowersCount int    `json:"followers_count"`
    FollowingCount int    `json:"following_count"`
    Posts         []Post  `json:"posts,omitempty"`       // empty or omitted if not allowed
    IsSelf        bool    `json:"is_self"`
    IsFollowing   bool    `json:"is_following"`
}

