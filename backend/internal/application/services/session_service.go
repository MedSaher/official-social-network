package services

import (
	"fmt"
	"time"

	"social_network/internal/application/utils"
	"social_network/internal/domain/ports/repository"
)

type SessionServiceImpl struct {
	sessionRepo repository.SessionRepository
	userRepo    repository.UserRepository
}

func NewSessionService(userRepo repository.UserRepository, sessionRepo repository.SessionRepository) *SessionServiceImpl {
	return &SessionServiceImpl{
		userRepo:    userRepo,
		sessionRepo: sessionRepo,
	}
}

func (s *SessionServiceImpl) CreateSession(userID int) (string, time.Time, error) {
	// Vérifier si l'utilisateur existe
	_, err := s.userRepo.GetUserByID(userID)
	if err != nil {
		return "", time.Time{}, fmt.Errorf("utilisateur introuvable: %w", err)
	}

	token, err := utils.GenerateRandomToken(32)
	if err != nil {
		return "", time.Time{}, fmt.Errorf("erreur token: %w", err)
	}

	// Expiration 24h
	expiresAt := time.Now().Add(24 * time.Hour)

	// Sauvegarde dans la base de données
	err = s.sessionRepo.CreateSession(userID, token, expiresAt)
	if err != nil {
		return "", time.Time{}, err
	}

	return token, expiresAt, nil
}

func (s *SessionServiceImpl) DestroySession(token string) error {
	return s.sessionRepo.DeleteSessionByToken(token)
}

func (s *SessionServiceImpl) IsValidSession(token string) bool {
	_, err := s.sessionRepo.GetSessionByToken(token)
	return err == nil
}

func (s *SessionServiceImpl) GetUserIdFromSession(token string) (int, error) {
	return s.sessionRepo.GetSessionByToken(token)
}
