package handlers

import (
	"encoding/json"
	"fmt"
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
		Email         string `json:"email"`
		Password      string `json:"password"`
		FirstName     string `json:"firstName"`
		LastName      string `json:"lastName"`
		Gender        string `json:"gender"`
		DateOfBirth   string `json:"dateOfBirth"`
		AboutMe       string `json:"aboutMe"`
		PrivacyStatus string `json:"privacyStatus"`
	}

	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		utils.ResponseJSON(w, http.StatusBadRequest, map[string]any{"error": "Invalid request body"})
		return
	}

	// Validate required fields
	if payload.Email == "" || payload.Password == "" || payload.FirstName == "" ||
		payload.LastName == "" || payload.DateOfBirth == "" || payload.Gender == "" || payload.PrivacyStatus == "" {
		utils.ResponseJSON(w, http.StatusBadRequest, map[string]any{"error": "Missing required fields"})
		return
	}

	// Validate gender
	if payload.Gender != "male" && payload.Gender != "female" {
		utils.ResponseJSON(w, http.StatusBadRequest, map[string]any{"error": "Invalid gender value"})
		return
	}

	// Validate privacy status
	validPrivacy := map[string]bool{
		"public":         true,
		"private":        true,
		"almost_private": true,
	}
	if !validPrivacy[payload.PrivacyStatus] {
		utils.ResponseJSON(w, http.StatusBadRequest, map[string]any{"error": "Invalid privacy status"})
		return
	}

	user := &models.User{
		Email:         payload.Email,
		Password:      payload.Password,
		FirstName:     payload.FirstName,
		LastName:      payload.LastName,
		Gender:        payload.Gender,
		DateOfBirth:   payload.DateOfBirth,
		AboutMe:       stringPtr(payload.AboutMe),
		PrivacyStatus: payload.PrivacyStatus,
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

	var payload struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		utils.ResponseJSON(w, http.StatusBadRequest, map[string]any{"error": "Invalid request body"})
		return
	}

	user, err := h.userService.Authenticate(payload.Email, payload.Password)
	if err != nil {
		utils.ResponseJSON(w, http.StatusUnauthorized, map[string]any{"error": "Authentication failed"})
		return
	}
	token, expiresAt, err := h.sessionService.CreateSession(user.Id)
	if err != nil {
		utils.ResponseJSON(w, http.StatusInternalServerError, map[string]any{"error": "Failed to create session"})
		return
	}
	fmt.Printf("token: %s, expiresAt: %s\n", token, expiresAt)
	http.SetCookie(w, &http.Cookie{
		Name:     "session_token",
		Value:    token,
		Expires:  expiresAt,
		SameSite: http.SameSiteLaxMode,
	})

	utils.ResponseJSON(w, http.StatusOK, map[string]any{
		"user_id":   user.Id,
		"token":     token,
		"expiresAt": expiresAt.Format(time.RFC3339),
	})
}

// ----------- Logout -----------
func (userHandler *UserHandler) Logout(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("session_token")
	if err != nil {
		utils.ResponseJSON(w, http.StatusUnauthorized, map[string]any{"message": "invalid token"})
		return
	}
	token := cookie.Value

	fmt.Println("log out: --------------> ", token)

	// userId, err := userHandler.sessionServ.GetUserIdFromSession(token)
	// if err != nil {
	// 	utils.ResponseJSON(w, http.StatusUnauthorized, map[string]any{"message": "invalid session"})
	// 	return
	// }

	err = userHandler.sessionService.DestroySession(token)
	if err != nil {
		utils.ResponseJSON(w, http.StatusInternalServerError, map[string]any{"message": "failed to logout"})
		return
	}

	// client := &services.Client{UserId: userId}
	// userHandler.chatBroker.Unregister <- client

	http.SetCookie(w, &http.Cookie{
		Name:     "session_token",
		Value:    "",
		Path:     "/",
		Expires:  time.Unix(0, 0),
		MaxAge:   -1,
		HttpOnly: true,
	})

	utils.ResponseJSON(w, http.StatusCreated, map[string]string{"message": "User logged out successfully"})
}

// Create a function to check if the user has a session:
func (userHandler *UserHandler) CheckSession(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		utils.ResponseJSON(w, http.StatusMethodNotAllowed, map[string]any{"error": "method not alowed"})
		return
	}
	cookie, err := r.Cookie("session_token")
	if err != nil {
		utils.ResponseJSON(w, http.StatusUnauthorized, map[string]any{"message": "invalid token"})
		return
	}

	token := cookie.Value
	logged := userHandler.sessionService.IsValidSession(token)
	if !logged {
		utils.ResponseJSON(w, http.StatusUnauthorized, map[string]any{"message": "invalid token"})
		return
	}

	utils.ResponseJSON(w, http.StatusOK, map[string]string{"message": "User has a valid session"})
}
