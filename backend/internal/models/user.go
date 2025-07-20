package models

// Create a structure to represent the user:

type User struct {
	Id           int     `json:"id"`
	Nickname     string  `json:"nickname"`
	Username     string  `json:"username"`
	DateOfBirth  string  `json:"dateOfBirth"` // use string or time.Time
	Gender       string  `json:"gender"`
	PasswordHash string  `json:"password"`    // stored hashed
	Email        string  `json:"email"`
	FirstName    string  `json:"firstName"`
	LastName     string  `json:"lastName"`
	AvatarPath   *string `json:"avatarUrl"`   // nullable
	AboutMe      *string `json:"aboutMe"`     // nullable
	IsPublic     bool    `json:"privacyStatus"`
	CreatedAt    string  `json:"createdAt"`   // optional, for info
}


// Create a model to ease working on chat
type ChatUser struct {
	Id          int    `json:"id"`
	NickName    string `json:"nick_name"`
	IsOnline    bool   `json:"is_online"`
	UnreadCount int    `json:"unread_count"`
}
