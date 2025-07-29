package repository

import (
	"social_network/internal/domain/models"
)

// Interface
type MessageRepository interface {
	InsertMessage(m *models.PrivateMessage) error
	GetChatHistory(client, guest int, offset int, limit int) ([]*models.PrivateMessage, error)
	GetLastMessage(user1ID, user2ID int) (*models.PrivateMessage, error)
	MarkMessagesAsRead(senderID, receiverID int) error
	GetUnreadMessageCount(userID int) (int, error)
	GetUnreadMessages(userID int) ([]*models.PrivateMessage, error)
}
