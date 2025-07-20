package db

import (
	"database/sql"
	"fmt"
	"time"

	"social_network/internal/domain/ports/repository"
)

// Implémentation de l'interface SessionRepository
type sessionRepositoryImpl struct {
	db *sql.DB
}

// NewSessionRepository crée une nouvelle instance de l'implémentation
func NewSessionRepository(db *sql.DB) repository.SessionRepository {
	return &sessionRepositoryImpl{db: db}
}

func (sr *sessionRepositoryImpl) CreateSession(userID int, token string, expiresAt time.Time) error {
	_, err := sr.db.Exec("DELETE FROM sessions WHERE user_id = ?", userID)
	if err != nil {
		return fmt.Errorf("error deleting existing sessions: %w", err)
	}

	_, err = sr.db.Exec(
		"INSERT INTO sessions (user_id, session_token, expires_at) VALUES (?, ?, ?)",
		userID, token, expiresAt,
	)
	if err != nil {
		return fmt.Errorf("error creating session: %w", err)
	}
	return nil
}

func (sr *sessionRepositoryImpl) DeleteSessionByToken(token string) error {
	_, err := sr.db.Exec("DELETE FROM sessions WHERE session_token = ?", token)
	return err
}

func (sr *sessionRepositoryImpl) GetSessionByToken(token string) (int, error) {
	query := `SELECT user_id FROM sessions WHERE session_token = ?`
	var userID int
	err := sr.db.QueryRow(query, token).Scan(&userID)
	if err != nil {
		return 0, err
	}
	return userID, nil
}
