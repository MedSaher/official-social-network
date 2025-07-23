package models
type User struct {
    Id            int
    Email         string
    Password      string   // hashed password
    FirstName     string
    LastName      string
    DateOfBirth   string
    AvatarPath    *string  // nullable avatar URL or path
    UserName      string
    AboutMe       *string  // nullable about me
    PrivacyStatus string   // "public", "private", or "almost_private"
    Gender        string   // "male" or "female"
    CreatedAt     string   // ideally time.Time, but string for now
}
