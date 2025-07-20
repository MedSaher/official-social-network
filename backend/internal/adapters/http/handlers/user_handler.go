package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"social_network/internal/adapters/http/utils"
	"social_network/internal/domain/models"
	"social_network/internal/domain/ports/service"
)

type UserHandler struct {
	userService    service.UserService
	sessionService service.SessionService
}

func NewUserHandler(userSvc service.UserService, sessionSvc service.SessionService) *UserHandler {
	return &UserHandler{
		userService:    userSvc,
		sessionService: sessionSvc,
	}
}

// Helper function to convert string to pointer
func stringPtr(s string) *string {
	if s == "" {
		return nil
	}
	return &s
}

// ----------- Register -----------
func (h *UserHandler) Register(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		utils.ResponseJSON(w, http.StatusMethodNotAllowed, map[string]any{"error": "Method not allowed"})
		return
	}

	var payload struct {
		NickName    string `json:"nickname"`
		UserName    string `json:"username"`
		Email       string `json:"email"`
		Password    string `json:"password"`
		FirstName   string `json:"firstName"`
		LastName    string `json:"lastName"`
		Gender      string `json:"gender"`
		DateOfBirth string `json:"dateOfBirth"`
		AboutMe     string `json:"aboutMe"`
	}

	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		utils.ResponseJSON(w, http.StatusBadRequest, map[string]any{"error": "Invalid request body"})
		return
	}

	user := &models.User{
		NickName:    payload.NickName,
		UserName:    payload.UserName,
		Email:       payload.Email,
		Password:    payload.Password,
		FirstName:   payload.FirstName,
		LastName:    payload.LastName,
		Gender:      payload.Gender,
		DateOfBirth: payload.DateOfBirth,
		AboutMe:     stringPtr(payload.AboutMe),
	}

	if err := h.userService.Register(user); err != nil {
		utils.ResponseJSON(w, http.StatusInternalServerError, map[string]any{"error": err.Error()})
		return
	}

	utils.ResponseJSON(w, http.StatusCreated, map[string]any{
		"success": true,
		"message": "User registered successfully. Please login.",
	})
}
// ----------- Login -----------
func (h *UserHandler) Login(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		utils.ResponseJSON(w, http.StatusMethodNotAllowed, map[string]any{"error": "Method not allowed"})
		return
	}

	var cred struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := json.NewDecoder(r.Body).Decode(&cred); err != nil {
		utils.ResponseJSON(w, http.StatusBadRequest, map[string]any{"error": "Invalid request body"})
		return
	}

	user, err := h.userService.Authenticate(cred.Email, cred.Password)
	if err != nil {
		utils.ResponseJSON(w, http.StatusUnauthorized, map[string]any{"error": "Authentication failed"})
		return
	}

	token, expiresAt, err := h.sessionService.CreateSession(user.Id)
	if err != nil {
		utils.ResponseJSON(w, http.StatusInternalServerError, map[string]any{"error": "Failed to create session"})
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:     "session_token",
		Value:    token,
		Expires:  expiresAt,
		HttpOnly: true,
		SameSite: http.SameSiteLaxMode,
	})

	utils.ResponseJSON(w, http.StatusOK, map[string]any{
		"user_id":   user.Id,
		"token":     token,
		"expiresAt": expiresAt.Format(time.RFC3339),
	})
}

// ----------- Logout -----------
func (h *UserHandler) Logout(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("session_token")
	if err != nil {
		utils.ResponseJSON(w, http.StatusUnauthorized, map[string]any{"error": "invalid session"})
		return
	}
	token := cookie.Value

	// Solution 1 : ignorer userId si tu t'en sers pas
	_, err = h.sessionService.GetUserIdFromSession(token)
	if err != nil {
		utils.ResponseJSON(w, http.StatusUnauthorized, map[string]any{"error": "session not found"})
		return
	}

	err = h.sessionService.DestroySession(token)
	if err != nil {
		utils.ResponseJSON(w, http.StatusInternalServerError, map[string]any{"error": "failed to destroy session"})
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:     "session_token",
		Value:    "",
		Path:     "/",
		MaxAge:   -1,
		Expires:  time.Unix(0, 0),
		HttpOnly: true,
	})

	utils.ResponseJSON(w, http.StatusOK, map[string]string{"message": "User logged out successfully"})
}
