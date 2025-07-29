package utils

import (
	"errors"
	"net/http"

	"social_network/internal/domain/ports/service"
)

func GetCurrentUserID(r *http.Request, sessionService service.SessionService) (int, error) {
	cookie, err := r.Cookie("session_token")
	if err != nil {
		return 0, errors.New("missing session token")
	}
	userID, err := sessionService.GetUserIdFromSession(cookie.Value)
	if err != nil {
		return 0, errors.New("invalid session")
	}

	return userID, nil
}
