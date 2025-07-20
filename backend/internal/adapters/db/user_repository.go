package db

import (
	"database/sql"

	"social_network/internal/domain/models"
	"social_network/internal/domain/ports/repository"
)

type userRepositoryImpl struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) repository.UserRepository {
	return &userRepositoryImpl{db: db}
}

func (r *userRepositoryImpl) RegisterNewUser(user *models.User) error {
	query := `
		INSERT INTO users (
			nickname, username, date_of_birth, gender, password_hash,
			email, first_name, last_name, avatar_path, about_me, is_public
		)
		VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
	`
	_, err := r.db.Exec(
		query,
		user.Nickname,
		user.Username,
		user.DateOfBirth,
		user.Gender,
		user.PasswordHash,
		user.Email,
		user.FirstName,
		user.LastName,
		user.AvatarPath,
		user.AboutMe,
		boolToInt(user.IsPublic),
	)
	return err
}

func (r *userRepositoryImpl) GetUserByID(id int) (*models.User, error) {
	query := `
		SELECT id, nickname, username, date_of_birth, gender, password_hash,
		       email, first_name, last_name, avatar_path, about_me, is_public, created_at
		FROM users WHERE id = ?
	`
	user := &models.User{}
	err := r.db.QueryRow(query, id).Scan(
		&user.Id,
		&user.Nickname,
		&user.Username,
		&user.DateOfBirth,
		&user.Gender,
		&user.PasswordHash,
		&user.Email,
		&user.FirstName,
		&user.LastName,
		&user.AvatarPath,
		&user.AboutMe,
		&user.IsPublic,
		&user.CreatedAt,
	)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (r *userRepositoryImpl) GetUserByEmail(email string) (*models.User, error) {
	query := `
		SELECT id, nickname, username, date_of_birth, gender, password_hash,
		       email, first_name, last_name, avatar_path, about_me, is_public, created_at
		FROM users WHERE email = ?
	`
	user := &models.User{}
	err := r.db.QueryRow(query, email).Scan(
		&user.Id,
		&user.Nickname,
		&user.Username,
		&user.DateOfBirth,
		&user.Gender,
		&user.PasswordHash,
		&user.Email,
		&user.FirstName,
		&user.LastName,
		&user.AvatarPath,
		&user.AboutMe,
		&user.IsPublic,
		&user.CreatedAt,
	)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func boolToInt(b bool) int {
	if b {
		return 1
	}
	return 0
}
