package repository

import "time"

// SessionRepository définit les méthodes accessibles au domaine
type SessionRepository interface {
	CreateSession(userID int, token string, expiresAt time.Time) error
	DeleteSessionByToken(token string) error
	GetSessionByToken(token string) (int, error)
}
