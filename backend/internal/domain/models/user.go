package models

type User struct {
	Id            int     `json:"id"`
	Email         string  `json:"email"`
	Password      string  `json:"password"`      
	FirstName     string  `json:"firstName"`
	LastName      string  `json:"lastName"`
	DateOfBirth   string  `json:"dateOfBirth"`   
	AvatarPath    *string `json:"avatarUrl"`     // nullable
	UserName      string  `json:"username"`
	AboutMe       *string `json:"aboutMe"`       // nullable
	PrivacyStatus string  `json:"privacyStatus"` // "public", "private", "almost_private"
	Gender        string  `json:"gender"`        
	CreatedAt     string  `json:"createdAt"`     // Ideally time.Time if you parse it
}
