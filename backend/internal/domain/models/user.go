package models

type User struct {
	Id          int   `json:"id"`
	NickName    string `json:"nickname"`
	Username    string `json:"username"`
	Age         int    `json:"age"`
	DateOfBirth string `json:"dateOfBirth"`
	Gender      string `json:"gender"`
	FirstName   string `json:"firstName"`
	LastName    string `json:"lastName"`
	Email       string `json:"email"`
	Password    string `json:"password"`
	About       string `json:"aboutMe"`
}