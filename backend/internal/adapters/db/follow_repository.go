package db

import (
	"database/sql"
	"fmt"
	"social_network/internal/domain/models"
	"social_network/internal/domain/ports/repository"
)

type FollowRepositoryImpl struct {
	db *sql.DB
}

func NewFollowRepository(db *sql.DB) repository.FollowRepository {
	return &FollowRepositoryImpl{db: db}
}

func (r *FollowRepositoryImpl) CreateFollow(follow *models.Follow) error {
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

func (r *FollowRepositoryImpl) AcceptFollow(followerID, followingID int) error {
	query := `
		UPDATE follows
		SET status = 'accepted'
		WHERE follower_id = ? AND following_id = ? AND status = 'pending'
	`

	result, err := r.db.Exec(query, followerID, followingID)
	if err != nil {
		return fmt.Errorf("error accepting follow request: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("error checking rows affected: %w", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("no pending follow request found for follower ID %d and following ID %d", followerID, followingID)
	}

	return nil
}

func (r *FollowRepositoryImpl) DeclineFollow(followerID, followingID int) error {
	query := `
		UPDATE follows
		SET status = 'declined'
		WHERE follower_id = ? AND following_id = ? AND status = 'pending'
	`

	result, err := r.db.Exec(query, followerID, followingID)
	if err != nil {
		return fmt.Errorf("error declining follow request: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("error checking rows affected: %w", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("no pending follow request found for follower ID %d and following ID %d", followerID, followingID)
	}

	return nil
}

func (r *FollowRepositoryImpl) DeleteFollow(followerID, followingID int) error {
	query := `
		DELETE FROM follows
		WHERE follower_id = ? AND following_id = ?
	`

	_, err := r.db.Exec(query, followerID, followingID)
	if err != nil {
		return fmt.Errorf("error deleting follow relationship: %w", err)
	}

	return nil
}

func (r *FollowRepositoryImpl) GetStatusFollow(followerID, followingID int) (string, error) {
	query := `
		SELECT status
		FROM follows
		WHERE follower_id = ? AND following_id = ?
	`

	var status string
	err := r.db.QueryRow(query, followerID, followingID).Scan(&status)
	if err != nil {
		if err == sql.ErrNoRows {
			return "", fmt.Errorf("no follow relationship found for follower ID %d and following ID %d", followerID, followingID)
		}
		return "", fmt.Errorf("error retrieving follow status: %w", err)
	}

	return status, nil
}

func (r *FollowRepositoryImpl) GetFollowers(userID int) ([]models.FollowerInfo, error) {
	query := `
		SELECT u.id, u.user_name, f.status
		FROM follows f
		JOIN users u ON f.follower_id = u.id
		WHERE f.following_id = ?
	`

	rows, err := r.db.Query(query, userID)
	if err != nil {
		return nil, fmt.Errorf("error retrieving followers: %w", err)
	}
	defer rows.Close()

	var followers []models.FollowerInfo
	for rows.Next() {
		var info models.FollowerInfo
		if err := rows.Scan(&info.UserID, &info.UserName, &info.Status); err != nil {
			return nil, fmt.Errorf("error scanning row: %w", err)
		}
		followers = append(followers, info)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating over rows: %w", err)
	}

	return followers, nil
}

func (r *FollowRepositoryImpl) GetFollowing(userID int) ([]models.FollowerInfo, error) {
	query := `
		SELECT u.id, u.user_name, f.status
		FROM follows f
		JOIN users u ON f.following_id = u.id
		WHERE f.follower_id = ?
	`

	rows, err := r.db.Query(query, userID)
	if err != nil {
		return nil, fmt.Errorf("error retrieving following: %w", err)
	}
	defer rows.Close()

	var following []models.FollowerInfo
	for rows.Next() {
		var info models.FollowerInfo
		if err := rows.Scan(&info.UserID, &info.UserName, &info.Status); err != nil {
			return nil, fmt.Errorf("error scanning row: %w", err)
		}
		following = append(following, info)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating over rows: %w", err)
	}

	return following, nil
}
