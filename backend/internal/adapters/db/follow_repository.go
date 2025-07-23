package db

import (
	"database/sql"
	"fmt"
	"social_network/internal/domain/models"
	"social_network/internal/domain/ports/repository"
)

type followRepositoryImpl struct {
	db *sql.DB
}

func NewFollowRepository(db *sql.DB) repository.FollowRepository {
	return &followRepositoryImpl{db: db}
}

func (r *followRepositoryImpl) CreateFollow(follow *models.Follow) error {
	query := `
		INSERT INTO follows (follower_id, following_id, status)
		VALUES (?, ?, ?)
	`

	_, err := r.db.Exec(query, follow.FollowerID, follow.FollowingID, follow.Status)
	if err != nil {
		return fmt.Errorf("error creating follow relationship: %w", err)
	}

	return nil
}
