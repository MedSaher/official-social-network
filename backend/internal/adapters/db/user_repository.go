package db

import (
	"database/sql"
	"fmt"

	"social_network/internal/domain/models"
	"social_network/internal/domain/ports/repository"
)

type UserRepositoryImpl struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) repository.UserRepository {
	return &UserRepositoryImpl{db: db}
}

func (r *UserRepositoryImpl) RegisterNewUser(user *models.User) error {
    query := `
        INSERT INTO users (
            email, password_hash, first_name, last_name, date_of_birth,
            avatar_path, user_name, about_me, privacy_status, gender
        )
        VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
    `

    _, err := r.db.Exec(
        query,
        user.Email,
        user.Password, 
        user.FirstName,
        user.LastName,
        user.DateOfBirth,
        user.AvatarPath, 
        user.UserName,     
        user.AboutMe,
        user.PrivacyStatus,
        user.Gender,       
    )
    return err
}

func (r *UserRepositoryImpl) GetUserByID(id int) (*models.User, error) {
		fmt.Println(" user by ID:", id)
	query := `
		SELECT id, user_name, date_of_birth, gender, password_hash,
		       email, first_name, last_name, avatar_path, about_me, privacy_status, created_at
		FROM users WHERE id = ?
	`
	user := &models.User{}
	err := r.db.QueryRow(query, id).Scan(
		&user.Id,
		&user.UserName,
		&user.DateOfBirth,
		&user.Gender,
		&user.Password,
		&user.Email,
		&user.FirstName,
		&user.LastName,
		&user.AvatarPath,
		&user.AboutMe,
		&user.PrivacyStatus,
		&user.CreatedAt,
	)
		fmt.Println("Fetching user by ID:", user.Id)

	if err != nil {
		return nil, err
	}
	return user, nil
}

func (r *UserRepositoryImpl) GetUserByEmail(email string) (*models.User, error) {
	query := `
		SELECT id, user_name, date_of_birth, gender, password_hash,
		       email, first_name, last_name, avatar_path, about_me, privacy_status, created_at
		FROM users WHERE email = ?
	`

	user := &models.User{}
	err := r.db.QueryRow(query, email).Scan(
		&user.Id,
		&user.UserName,
		&user.DateOfBirth,
		&user.Gender,
		&user.Password,      
		&user.Email,
		&user.FirstName,
		&user.LastName,
		&user.AvatarPath,
		&user.AboutMe,
		&user.PrivacyStatus,
		&user.CreatedAt,
	)
	if err != nil {
		return nil, err
	}

	return user, nil
}

