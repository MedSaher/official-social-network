package db

import (
	"context"
	"database/sql"
	"fmt"
	"social_network/internal/domain/models"
	"social_network/internal/domain/ports/repository"
	"time"
)

type PostRepository struct {
	db *sql.DB
}

func NewPostRepository(db *sql.DB) repository.PostRepository {
	return &PostRepository{db: db}
}

// CreatePost inserts a new post into the DB and returns the created post with its auto-generated fields.
func (r *PostRepository) CreatePost(ctx context.Context, userID int, groupID *int, content, privacy, imagePath string) (models.Post, error) {
	var post models.Post

	// Validate privacy before insertion (optional but recommended)
	validPrivacy := map[string]bool{"public": true, "almost_private": true, "private": true}
	if !validPrivacy[privacy] {
		privacy = "public" // default fallback
	}

	// SQLite Insert query with RETURNING clause to get inserted row values
	query := `
        INSERT INTO posts (user_id, group_id, content, image_path, privacy, created_at, updated_at)
        VALUES (?, ?, ?, ?, ?, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP)
        RETURNING id, user_id, group_id, content, image_path, privacy, created_at, updated_at;
    `

	// Use sql.NullInt64 / sql.NullString for nullable fields if scanning into Post struct directly
	// But here, we'll scan directly into pointers for group_id and image_path

	row := r.db.QueryRowContext(ctx, query,
		userID,
		groupID,
		content,
		sql.NullString{String: imagePath, Valid: imagePath != ""},
		privacy,
	)

	var groupIDNull sql.NullInt64
	var imagePathNull sql.NullString
	var createdAtStr, updatedAtStr string

	err := row.Scan(
		&post.ID,       // id
		&post.UserID,   // user_id
		&groupIDNull,   // group_id
		&post.Content,  // content ← was missing
		&imagePathNull, // image_path
		&post.Privacy,  // privacy
		&createdAtStr,  // created_at
		&updatedAtStr,  // updated_at
	)
	if err != nil {
		return models.Post{}, err
	}

	// Handle nullable group_id
	if groupIDNull.Valid {
		val := int(groupIDNull.Int64)
		post.GroupID = &val
	}

	// Handle nullable image_path
	if imagePathNull.Valid {
		post.ImagePath = &imagePathNull.String
	}

	// Parse timestamps (SQLite returns DATETIME as string)
	post.CreatedAt, err = time.Parse(time.RFC3339, createdAtStr)
	if err != nil {
		return models.Post{}, fmt.Errorf("invalid created_at timestamp format: %v", createdAtStr)
	}

	post.UpdatedAt, err = time.Parse(time.RFC3339, updatedAtStr)
	if err != nil {
		return models.Post{}, fmt.Errorf("invalid updated_at timestamp format: %v", updatedAtStr)
	}

	post.Content = content

	return post, nil
}
