package repositories

import (
	"database/sql"
	"fmt"

	"social_network/internal/models"
)

// Create an interface to represent all the user repositoru functionalities:
type UsersRepositoryLayer interface {
	RegisterNewUser(user *models.User) error
	GetUserByEmail(email string) (*models.User, error)
	GetUserByID(id int) (*models.User, error)
	GetSortedUsersForChat(myID, offset, limit int) ([]*models.ChatUser, error)
}

// Create a structure to represent to implemente the contract with the repo interface:
type UsersRepository struct {
	db *sql.DB
}

// Create a new instance from the user crepository structure:
func NewUsersRepository(database *sql.DB) *UsersRepository {
	userRepo := new(UsersRepository)
	userRepo.db = database
	// return &UsersRepository{db: database}
	return userRepo
}

// Create a function to register a new user:
func (userRepo *UsersRepository) RegisterNewUser(user *models.User) error {
	query := `
		INSERT INTO users (
			nickname, username, date_of_birth, gender, password_hash,
			email, first_name, last_name, avatar_path, about_me, is_public
		)
		VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
	`
	_, err := userRepo.db.Exec(
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

func boolToInt(b bool) int {
	if b {
		return 1
	}
	return 0
}


// get user bu id:
func (userRepo *UsersRepository) GetUserByID(id int) (*models.User, error) {
	query := `
		SELECT id, nickname, username, date_of_birth, gender, password_hash,
		       email, first_name, last_name, avatar_path, about_me, is_public, created_at
		FROM users WHERE id = ?
	`
	user := &models.User{}
	err := userRepo.db.QueryRow(query, id).Scan(
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


// Get all users:
func (userRepo *UsersRepository) GetSortedUsersForChat(myID, offset, limit int) ([]*models.ChatUser, error) {
	query := `
    SELECT id, nick_name, unread_count
    FROM (
        -- Users who have chatted with me
        SELECT 
            u.id, 
            u.nick_name, 
            MAX(pm.created_at) AS last_message_time,
            (
                SELECT COUNT(*) 
                FROM private_messages 
                WHERE sender_id = u.id AND receiver_id = ? AND is_read = 0
            ) AS unread_count
        FROM users u
        JOIN private_messages pm
            ON (u.id = pm.sender_id AND pm.receiver_id = ?) 
            OR (u.id = pm.receiver_id AND pm.sender_id = ?)
        WHERE u.id != ?
        GROUP BY u.id

        UNION ALL

        -- Users who have NOT chatted with me
        SELECT 
            u.id, 
            u.nick_name, 
            NULL as last_message_time,
            0 as unread_count
        FROM users u
        WHERE u.id != ? AND u.id NOT IN (
            SELECT 
                CASE 
                    WHEN pm.sender_id = ? THEN pm.receiver_id
                    ELSE pm.sender_id
                END
            FROM private_messages pm
            WHERE pm.sender_id = ? OR pm.receiver_id = ?
        )
    ) AS all_users
    ORDER BY 
        last_message_time IS NULL,       
        last_message_time DESC,        
        LOWER(nick_name) ASC           
    LIMIT ? OFFSET ?;
    `

	rows, err := userRepo.db.Query(
		query,
		myID, myID, myID, myID, // For first subquery (unread count, joins, filter)
		myID, myID, myID, myID, // For second subquery (not-in logic)
		limit, offset,
	)
	if err != nil {
		return nil, fmt.Errorf("query error: %v", err)
	}
	defer rows.Close()

	var users []*models.ChatUser
	for rows.Next() {
		chatUser := &models.ChatUser{}
		if err := rows.Scan(&chatUser.Id, &chatUser.NickName, &chatUser.UnreadCount); err != nil {
			return nil, err
		}

		// For now, you can determine online status elsewhere (e.g., session map)
		chatUser.IsOnline = false

		users = append(users, chatUser)
	}

	return users, nil
}

func (userRepo *UsersRepository) GetUserByEmail(email string) (*models.User, error) {
	query := "SELECT id, nickname, username, date_of_birth, gender, password_hash, email, first_name, last_name, avatar_path, about_me, is_public, created_at FROM users WHERE email = ?"

	user := &models.User{}
	err := userRepo.db.QueryRow(query, email).Scan(
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
