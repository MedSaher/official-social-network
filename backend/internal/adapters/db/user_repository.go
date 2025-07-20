package db

import (
	"database/sql"
	"fmt"

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
			nick_name, user_name, date_of_birth, gender, password_hash,
			email, first_name, last_name, avatar_path, about_me, is_public
		)
		VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
	`
	_, err := r.db.Exec(
		query,
		user.NickName,
		user.UserName,
		user.DateOfBirth,
		user.Gender,
		user.Password,
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
		fmt.Println(" user by ID:", id)
	query := `
		SELECT id, nick_name, user_name, date_of_birth, gender, password_hash,
		       email, first_name, last_name, avatar_path, about_me, is_public, created_at
		FROM users WHERE id = ?
	`
	user := &models.User{}
	err := r.db.QueryRow(query, id).Scan(
		&user.Id,
		&user.NickName,
		&user.UserName,
		&user.DateOfBirth,
		&user.Gender,
		&user.Password,
		&user.Email,
		&user.FirstName,
		&user.LastName,
		&user.AvatarPath,
		&user.AboutMe,
		&user.IsPublic,
		&user.CreatedAt,
	)
		fmt.Println("Fetching user by ID:", user.Id)

	if err != nil {
		return nil, err
	}
	return user, nil
}

func (r *userRepositoryImpl) GetUserByEmail(email string) (*models.User, error) {
	query := `
		SELECT id, nick_name, user_name, date_of_birth, gender, password_hash,
		       email, first_name, last_name, avatar_path, about_me, is_public, created_at
		FROM users WHERE email = ?
	`
	user := &models.User{}
	err := r.db.QueryRow(query, email).Scan(
		&user.Id,
		&user.NickName,
		&user.UserName,
		&user.DateOfBirth,
		&user.Gender,
		&user.Password,
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
