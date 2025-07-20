package services

import (
	"fmt"
	"social_network/internal/models"
	"social_network/internal/repositories"
)

// Create messages' service interface:
type MessagesServiceLayer interface {
	GetChatHistoryService(id int, sessionValue string, offset int, limit int) ([]*models.Message, error)
	MarkMessageAsRead(fromID, userId int) error
}

// Create the employee that will execute the message sevice interface:
type MessagesService struct {
	messageRepo repositories.MessageRepositoryLayer
	sessRepo    repositories.SessionsRepositoryLayer
}

// Instantiate the message service:
func NewMessageService(messRep *repositories.MessageRepository, sessionRepo *repositories.SessionsRepository) *MessagesService {
	return &MessagesService{
		messageRepo: messRep,
		sessRepo:    sessionRepo,
	}
}

// Get all the messages between the client and the chosen user:
func (mesSer *MessagesService) GetChatHistoryService(id int, sessionValue string, offset int, limit int) ([]*models.Message, error) {
	// Get client ID from session token
	clientId, err := mesSer.sessRepo.GetSessionByToken(sessionValue)
	if err!=nil {
		return nil, fmt.Errorf("invalid or expired session token")
	}
	// Retrieve chat history between client and selected user
	return mesSer.messageRepo.GetChatHistory(clientId, id, offset, limit)
}

// mark message as read service:
func (mesSer *MessagesService)MarkMessageAsRead(fromID, userId int) error {
	err := mesSer.messageRepo.MarkMessagesAsRead(fromID, userId)
	if err != nil {
		return err
	}
	return nil
}