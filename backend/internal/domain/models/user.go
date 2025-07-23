package models

type User struct {
	Id          int     `json:"id"`
	NickName    string  `json:"nickname"`
	UserName    string  `json:"username"`
	DateOfBirth string  `json:"dateOfBirth"` // use string or time.Time
	Gender      string  `json:"gender"`
	Password    string  `json:"password"` // stored hashed
	Email       string  `json:"email"`
	FirstName   string  `json:"firstName"`
	LastName    string  `json:"lastName"`
	AvatarPath  *string `json:"avatarUrl"` // nullable
	AboutMe     *string `json:"aboutMe"`   // nullable
	IsPublic    bool    `json:"is_public"`
	CreatedAt   string  `json:"createdAt"` // optional, for info
}
