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

type UserProfileDTO struct {
	Id            int     `json:"id"`
	UserName      string  `json:"username"`
	FirstName     string  `json:"firstName"`
	LastName      string  `json:"lastName"`
	AvatarUrl     *string `json:"avatarUrl"`
	Email         string  `json:"email"`
	AboutMe       *string `json:"aboutMe"`
	PrivacyStatus string  `json:"privacyStatus"`
	Gender        string  `json:"gender"`
	CreatedAt     string  `json:"createdAt"`
}

func UserProfileDTOFromUser(u *User) UserProfileDTO {
	return UserProfileDTO{
		Id:            u.Id,
		UserName:      u.UserName,
		FirstName:     u.FirstName,
		LastName:      u.LastName,
		AvatarUrl:     u.AvatarPath,
		Email:         u.Email,
		AboutMe:       u.AboutMe,
		PrivacyStatus: u.PrivacyStatus,
		Gender:        u.Gender,
		CreatedAt:     u.CreatedAt,
	}
}