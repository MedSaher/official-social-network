package models


type User struct {
	Id            int
	Email         string
	Password      string   // This stores the password_hash from DB
	FirstName     string
	LastName      string
	DateOfBirth   string   // Ideally time.Time if you parse the date
	AvatarPath    *string  // nullable
	UserName      string
	AboutMe       *string  // nullable
	PrivacyStatus string   // should be "public", "private", or "almost_private"
	Gender        string   // should be "male" or "female"
	CreatedAt     string   // Ideally time.Time if you parse it
}
