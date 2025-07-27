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

// repository/group_repository.go
func (r *GroupRepository) GetAllGroups(ctx context.Context) ([]models.Group, error) {
	rows, err := r.db.QueryContext(ctx, `
		SELECT id, creator_id, title, description, created_at, updated_at
		FROM groups
		ORDER BY created_at DESC
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var groups []models.Group
	for rows.Next() {
		var g models.Group
		if err := rows.Scan(&g.ID, &g.CreatorID, &g.Title, &g.Description, &g.CreatedAt, &g.UpdatedAt); err != nil {
			return nil, err
		}
		groups = append(groups, g)
	}

	return groups, nil
}
