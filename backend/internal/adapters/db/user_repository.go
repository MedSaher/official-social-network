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
	_, err := r.db.Exec(`
		INSERT INTO users (
			nickname, username, date_of_birth, gender, password_hash, email, first_name, last_name, about_me
		) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)
	`, user.NickName, user.Username, user.DateOfBirth, user.Gender, user.Password,
		user.Email, user.FirstName, user.LastName, user.About)
	if err != nil {
		return fmt.Errorf("failed to register user: %w", err)
	}
	return nil
}
