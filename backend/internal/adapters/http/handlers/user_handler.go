package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
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

// Helper to convert string to *string
func stringPtr(s string) *string {
	if s == "" {
		return nil
	}
	return &s
}

func (h *UserHandler) Register(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Register handler called")
	if r.Method != http.MethodPost {
		utils.ResponseJSON(w, http.StatusMethodNotAllowed, map[string]any{"error": "Method not allowed"})
		return
	}

	if err := r.ParseMultipartForm(10 << 20); err != nil {
		utils.ResponseJSON(w, http.StatusBadRequest, map[string]any{"error": "Failed to parse form data"})
		return
	}

	// Extract form values first (needed for nickname)
	email := r.FormValue("email")
	password := r.FormValue("password")
	firstName := r.FormValue("firstName")
	lastName := r.FormValue("lastName")
	gender := r.FormValue("gender")
	dateOfBirth := r.FormValue("dateOfBirth")
	aboutMe := r.FormValue("aboutMe")
	privacyStatus := r.FormValue("privacyStatus")
	nickname := r.FormValue("nickname")

	// Prepare avatar file (if provided)
	var avatarPath *string
	file, header, err := r.FormFile("avatar")
	if err == nil {
		defer file.Close()

		ext := filepath.Ext(header.Filename)
		timestamp := time.Now().UnixNano()
		newFileName := fmt.Sprintf("%d%s", timestamp, ext)

		const uploadDir = "./avatars"
		if err := os.MkdirAll(uploadDir, os.ModePerm); err != nil {
			utils.ResponseJSON(w, http.StatusInternalServerError, map[string]any{"error": "Failed to create upload directory"})
			return
		}

		savePath := filepath.Join(uploadDir, newFileName)
		out, err := os.Create(savePath)
		if err != nil {
			utils.ResponseJSON(w, http.StatusInternalServerError, map[string]any{"error": "Failed to save avatar file"})
			return
		}
		defer out.Close()

		if _, err := io.Copy(out, file); err != nil {
			utils.ResponseJSON(w, http.StatusInternalServerError, map[string]any{"error": "Failed to write avatar file"})
			return
		}

		avatarPath = &savePath
	} else if err != http.ErrMissingFile {
		utils.ResponseJSON(w, http.StatusBadRequest, map[string]any{"error": "Error retrieving avatar file"})
		return
	}

	user := &models.User{
		Email:         email,
		Password:      password,
		FirstName:     firstName,
		LastName:      lastName,
		Gender:        gender,
		DateOfBirth:   dateOfBirth,
		AboutMe:       stringPtr(aboutMe),
		AvatarPath:    avatarPath,
		PrivacyStatus: privacyStatus,
		UserName:      nickname,
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

	utils.ResponseJSON(w, http.StatusOK, map[string]string{"success": "true"})
}

// Create a function to check if the user has a session:
func (h *UserHandler) CheckSession(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		utils.ResponseJSON(w, http.StatusMethodNotAllowed, map[string]any{"error": "method not alowed"})
		return
	}

	// Check if the user has a valid session
	_, err := utils.GetCurrentUserID(r, h.sessionService)
	if err != nil {
		utils.ResponseJSON(w, http.StatusUnauthorized, map[string]any{"message": "invalid token"})
		return
	}

	utils.ResponseJSON(w, http.StatusOK, map[string]string{"message": "User has a valid session"})
}

func (h *UserHandler) GetFullProfile(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		utils.ResponseJSON(w, http.StatusMethodNotAllowed, map[string]any{"error": "method not allowed"})
		return
	}
	loggedUserID, err := utils.GetCurrentUserID(r, h.sessionService)
	if err != nil {
		utils.ResponseJSON(w, http.StatusUnauthorized, map[string]any{"error": "unauthorized"})
		return
	}

	// Optionally get profile owner ID from query param if viewing another user
	profileID := r.URL.Query().Get("user_id")
	if profileID == "" {
		profileID = strconv.Itoa(loggedUserID)
	}
	profileOwnerID, err := strconv.Atoi(profileID)
	if err != nil {
		utils.ResponseJSON(w, http.StatusBadRequest, map[string]any{"error": "invalid user id"})
		return
	}

	profileData, err := h.userService.GetFullProfileData(loggedUserID, profileOwnerID)
	if err != nil {
		utils.ResponseJSON(w, http.StatusInternalServerError, map[string]any{"error": err.Error()})
		return
	}

	utils.ResponseJSON(w, http.StatusOK, profileData)
}

func (h *UserHandler) SearchUsers(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		utils.ResponseJSON(w, http.StatusMethodNotAllowed, map[string]any{"error": "Method not allowed"})
		return
	}

	query := r.URL.Query().Get("q")
	if len(query) < 2 {
		utils.ResponseJSON(w, http.StatusBadRequest, map[string]any{"error": "Query too short"})
		return
	}

	users, err := h.userService.SearchUsers(query)
	if err != nil {
		utils.ResponseJSON(w, http.StatusInternalServerError, map[string]any{"error": err.Error()})
		return
	}

	utils.ResponseJSON(w, http.StatusOK, users)
}

func (h *UserHandler) GetUserProfileByUsername(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		utils.ResponseJSON(w, http.StatusMethodNotAllowed, map[string]any{"error": "Method not allowed"})
		return
	}

	username := r.URL.Query().Get("username")
	if username == "" {
		utils.ResponseJSON(w, http.StatusBadRequest, map[string]any{"error": "username is required"})
		return
	}

	user, err := h.userService.GetUserProfileByUsername(username)
	if err != nil {
		utils.ResponseJSON(w, http.StatusInternalServerError, map[string]any{"error": err.Error()})
		return
	}
	if user == nil {
		utils.ResponseJSON(w, http.StatusNotFound, map[string]any{"error": "User not found"})
		return
	}

	utils.ResponseJSON(w, http.StatusOK, user)
}