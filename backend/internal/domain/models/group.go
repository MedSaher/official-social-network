package models

import "time"

type Group struct {
	ID          int       `json:"id"`
	CreatorID   int       `json:"creator_id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type GroupWithUserFlags struct {
	ID          int       `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	CreatorID   int       `json:"creator_id"`
	CreatedAt   time.Time `json:"created_at"`
	IsCreator bool `json:"is_creator"`
	IsMember  bool `json:"is_member"`
}

type GroupJoinRequest struct {
	ID        int    `json:"id"`
	UserID    int    `json:"user_id"`
	UserName  string `json:"user_name"`
	CreatedAt string `json:"created_at"`
}

type GroupMember struct {
	ID       int
	GroupID  int
	UserID   int
	Status   string
	Role     string
}

type GroupPost struct {
	ID        int       `json:"id"`
	Content   string    `json:"content"`
	CreatedAt string    `json:"created_at"`
	ImagePath *string   `json:"image_path,omitempty"`
	UserID    int       `json:"user_id"`
	UserName  string    `json:"user_name"`
}

type GroupInfo struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	CreatorID   int    `json:"creator_id"`
	PostsCount  int    `json:"posts_count"`
	EventsCount int    `json:"events_count"`
}

type GroupEvent struct {
	ID          int    `json:"id"`
	GroupID     int    `json:"group_id"`
	CreatorID   int    `json:"creator_id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	EventDate   string `json:"event_date"` // ISO string
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}
