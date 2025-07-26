package db

import (
	"context"
	"database/sql"
	"social_network/internal/domain/models"
)

type GroupRepository struct {
	db *sql.DB
}

func NewGroupRepository(db *sql.DB) *GroupRepository {
	return &GroupRepository{db: db}
}

func (r *GroupRepository) CreateGroup(ctx context.Context, g *models.Group) error {
	result, err := r.db.ExecContext(ctx, `
		INSERT INTO groups (creator_id, title, description)
		VALUES (?, ?, ?)`,
		g.CreatorID, g.Title, g.Description,
	)
	if err != nil {
		return err
	}

	lastID, err := result.LastInsertId()
	if err != nil {
		return err
	}
	g.ID = int(lastID)
	return nil
}
