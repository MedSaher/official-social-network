package models

import "time"

type Post struct {
    ID        int        `json:"id"`
    UserID    int        `json:"user_id"`
    GroupID   *int       `json:"group_id,omitempty"`  // nullable field → pointer to int
    Content   string     `json:"content"`
    ImagePath *string    `json:"image_path,omitempty"` // nullable field → pointer to string
    Privacy   string     `json:"privacy"`             // expect "public", "almost_private", or "private"
    CreatedAt time.Time  `json:"created_at"`
    UpdatedAt time.Time  `json:"updated_at"`
}

type Comment struct {
    ID        int       `json:"id"`
    PostID    int       `json:"post_id"`
    UserID    int       `json:"user_id"`
    Content   string    `json:"content"`
    CreatedAt time.Time `json:"created_at"`
}