package service

import "time"

type SessionService interface {
	CreateSession(userID int) (string, time.Time, error)
	DestroySession(token string) error
	// IsValidSession(token string) bool
	GetUserIdFromSession(token string) (int, error)
}