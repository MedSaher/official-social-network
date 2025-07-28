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

func (r *PostRepository) GetAllPosts(ctx context.Context) ([]models.Post, error) {
    rows, err := r.db.QueryContext(ctx, `
        SELECT id, user_id, group_id, content, image_path, privacy, created_at, updated_at
        FROM posts
        ORDER BY created_at DESC
    `)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var posts []models.Post

    for rows.Next() {
        var post models.Post
        var groupID sql.NullInt64
        var imagePath sql.NullString

        err := rows.Scan(
            &post.ID,
            &post.UserID,
            &groupID,
            &post.Content,
            &imagePath,
            &post.Privacy,
            &post.CreatedAt,
            &post.UpdatedAt,
        )
        if err != nil {
            return nil, err
        }

        if groupID.Valid {
            gid := int(groupID.Int64)
            post.GroupID = &gid
        }

        if imagePath.Valid {
            img := imagePath.String
            post.ImagePath = &img
        }

        posts = append(posts, post)
    }

    return posts, nil
}

func (r *PostRepository) GetCommentsByPostID(ctx context.Context, postID int) ([]models.Comment, error) {
    rows, err := r.db.QueryContext(ctx, `
        SELECT id, post_id, user_id, content, created_at
        FROM comments
        WHERE post_id = ?
        ORDER BY created_at ASC
    `, postID)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var comments []models.Comment
    for rows.Next() {
        var c models.Comment
        if err := rows.Scan(&c.ID, &c.PostID, &c.UserID, &c.Content, &c.CreatedAt); err != nil {
            return nil, err
        }
        comments = append(comments, c)
    }

    return comments, nil
}

func (r *PostRepository) CreateComment(ctx context.Context, c *models.Comment) error {
    result, err := r.db.ExecContext(ctx, `
        INSERT INTO comments (post_id, user_id, content)
        VALUES (?, ?, ?)`,
        c.PostID, c.UserID, c.Content,
    )
    if err != nil {
        return err
    }

    lastID, err := result.LastInsertId()
    if err != nil {
        return err
    }

    c.ID = int(lastID)
    // created_at is set by DB default, you may fetch it later if needed
    return nil
}

func (r *PostRepository) GetPostsByUserID(userID int) ([]models.Post, error) {

	query := `SELECT id, user_id, group_id, content, image_path, privacy, created_at, updated_at FROM posts WHERE user_id = ? ORDER BY created_at DESC`
	rows, err := r.db.Query(query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var posts []models.Post
	for rows.Next() {
		var p models.Post
		err := rows.Scan(
			&p.ID,
			&p.UserID,
			&p.GroupID,
			&p.Content,
			&p.ImagePath,
			&p.Privacy,
			&p.CreatedAt,
			&p.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		posts = append(posts, p)
	}
	return posts, nil
}
