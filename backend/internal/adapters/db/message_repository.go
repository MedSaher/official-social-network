package db

import (
	"database/sql"
	"fmt"
	"social_network/internal/domain/ports/repository"
	"social_network/internal/domain/models"
)


// Struct implementation - unexported (starts with lowercase)
type messageRepository struct {
	db *sql.DB
}

// Constructor returns the interface type but implemented by the struct pointer
func NewMessageRepository(db *sql.DB) repository.MessageRepository {
	return &messageRepository{db: db}
}

// Implement the methods with pointer receiver to the struct (messageRepository)
func (repo *messageRepository) InsertMessage(msg *models.PrivateMessage) error {
	query := `INSERT INTO private_messages(content, sender_id, receiver_id, is_read) VALUES(?, ?, ?, ?)`
	_, err := repo.db.Exec(query, msg.Content, msg.SenderID, msg.ReceiverID, msg.IsRead)
	return err
}

func (repo *messageRepository) GetChatHistory(client, guest, offset, limit int) ([]*models.PrivateMessage, error) {
	query := `
	SELECT id, content, sender_id, receiver_id, is_read, created_at
	FROM private_messages
	WHERE (sender_id = ? AND receiver_id = ?) OR (sender_id = ? AND receiver_id = ?)
	ORDER BY id DESC
	LIMIT ? OFFSET ?
	`

	rows, err := repo.db.Query(query, client, guest, guest, client, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var messages []*models.PrivateMessage
	for rows.Next() {
		msg := &models.PrivateMessage{}
		if err := rows.Scan(&msg.ID, &msg.Content, &msg.SenderID, &msg.ReceiverID, &msg.IsRead, &msg.CreatedAt); err != nil {
			return nil, err
		}
		messages = append(messages, msg)
	}

	return messages, nil
}

func (repo *messageRepository) GetLastMessage(client, guest int) (*models.PrivateMessage, error) {
	query := `
	SELECT id, content, sender_id, receiver_id, is_read, created_at
	FROM private_messages
	WHERE (sender_id = ? AND receiver_id = ?) OR (sender_id = ? AND receiver_id = ?)
	ORDER BY created_at DESC
	LIMIT 1;
	`

	row := repo.db.QueryRow(query, client, guest, guest, client)

	msg := &models.PrivateMessage{}
	err := row.Scan(&msg.ID, &msg.Content, &msg.SenderID, &msg.ReceiverID, &msg.IsRead, &msg.CreatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return msg, nil
}

func (repo *messageRepository) MarkMessagesAsRead(senderID, receiverID int) error {
	query := `
	UPDATE private_messages 
	SET is_read = 1 
	WHERE sender_id = ? AND receiver_id = ? AND is_read = 0
	`
	_, err := repo.db.Exec(query, senderID, receiverID)
	if err != nil {
		return fmt.Errorf("failed to mark messages as read: %w", err)
	}
	return nil
}

func (repo *messageRepository) GetUnreadMessageCount(userID int) (int, error) {
	query := `
	SELECT COUNT(*) 
	FROM private_messages 
	WHERE receiver_id = ? AND is_read = 0
	`
	var count int
	err := repo.db.QueryRow(query, userID).Scan(&count)
	if err != nil {
		return 0, fmt.Errorf("failed to get unread message count: %w", err)
	}
	return count, nil
}

func (repo *messageRepository) GetUnreadMessages(userID int) ([]*models.PrivateMessage, error) {
	query := `
	SELECT id, content, sender_id, receiver_id, is_read, created_at 
	FROM private_messages 
	WHERE receiver_id = ? AND is_read = 0 
	ORDER BY created_at ASC
	`

	rows, err := repo.db.Query(query, userID)
	if err != nil {
		return nil, fmt.Errorf("failed to get unread messages: %w", err)
	}
	defer rows.Close()

	var messages []*models.PrivateMessage
	for rows.Next() {
		msg := &models.PrivateMessage{}
		if err := rows.Scan(&msg.ID, &msg.Content, &msg.SenderID, &msg.ReceiverID, &msg.IsRead, &msg.CreatedAt); err != nil {
			return nil, err
		}
		messages = append(messages, msg)
	}

	return messages, nil
}
