package db

import (
	"context"
	"database/sql"
	"errors"

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

func (r *GroupRepository) GetAllGroupsForUser(ctx context.Context, userID int) ([]models.GroupWithUserFlags, error) {
	rows, err := r.db.QueryContext(ctx, `
		SELECT
			g.id,
			g.title,
			g.description,
			g.creator_id,
			g.created_at,
			g.updated_at,
			CASE WHEN g.creator_id = ? THEN 1 ELSE 0 END as is_creator,
			EXISTS (
				SELECT 1 FROM group_members gm
				WHERE gm.group_id = g.id AND gm.user_id = ? AND gm.status = 'accepted'
			) as is_member
		FROM groups g
		ORDER BY g.created_at DESC
	`, userID, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var groups []models.GroupWithUserFlags
	for rows.Next() {
		var g models.GroupWithUserFlags
		var isCreatorInt, isMemberInt int
		if err := rows.Scan(
			&g.ID, &g.Title, &g.Description, &g.CreatorID,
			&g.CreatedAt, &g.UpdatedAt,
			&isCreatorInt, &isMemberInt,
		); err != nil {
			return nil, err
		}
		g.IsCreator = isCreatorInt == 1
		g.IsMember = isMemberInt == 1
		groups = append(groups, g)
	}
	return groups, nil
}

func (r *GroupRepository) IsAlreadyMember(ctx context.Context, groupID int, userID int) (bool, error) {
	var exists int
	err := r.db.QueryRowContext(ctx, `
		SELECT 1 FROM group_members 
		WHERE group_id = ? AND user_id = ?
	`, groupID, userID).Scan(&exists)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return false, nil
		}
		return false, err
	}
	return true, nil
}

func (r *GroupRepository) CreateJoinRequest(ctx context.Context, groupID int, userID int) error {
	_, err := r.db.ExecContext(ctx, `
		INSERT INTO group_members (group_id, user_id, status, role)
		VALUES (?, ?, 'pending', 'member')
	`, groupID, userID)
	return err
}
