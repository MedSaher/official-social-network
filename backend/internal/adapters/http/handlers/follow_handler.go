package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"social_network/internal/adapters/http/utils"
	"social_network/internal/domain/models"
	"social_network/internal/domain/ports/service"
)

type FollowHandler struct {
	followService  service.FollowService
	sessionService service.SessionService
}

func NewFollowHandler(followSvc service.FollowService, sessionSvc service.SessionService) *FollowHandler {
	return &FollowHandler{
		followService:  followSvc,
		sessionService: sessionSvc,
	}
}

func (h *FollowHandler) CreateFollow(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		utils.ResponseJSON(w, http.StatusMethodNotAllowed, map[string]any{"error": "Method not allowed"})
		return
	}

	var payload struct {
		FollowingID int `json:"following_id"`
	}

	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		utils.ResponseJSON(w, http.StatusBadRequest, map[string]any{"error": "Invalid request body"})
		return
	}

	// Get the current user's ID from the session
	followerID, err := utils.GetCurrentUserID(r, h.sessionService)
	if err != nil {
		utils.ResponseJSON(w, http.StatusUnauthorized, map[string]any{"error": "Unauthorized"})
		return
	}

	follow := &models.Follow{
		FollowerID:  followerID,
		FollowingID: payload.FollowingID,
	}

	if err := h.followService.CreateFollow(follow); err != nil {
		utils.ResponseJSON(w, http.StatusInternalServerError, map[string]any{"error": err.Error()})
		return
	}

	utils.ResponseJSON(w, http.StatusCreated, map[string]any{
		"success": true,
		"message": "Follow request created successfully.",
	})
}

func (h *FollowHandler) AcceptFollow(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		utils.ResponseJSON(w, http.StatusMethodNotAllowed, map[string]any{"error": "Method not allowed"})
		return
	}

	var payload struct {
		FollowerID  int `json:"follower_id"`
		FollowingID int `json:"following_id"`
	}

	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		utils.ResponseJSON(w, http.StatusBadRequest, map[string]any{"error": "Invalid request body"})
		return
	}

	currentUserID, err := utils.GetCurrentUserID(r, h.sessionService)
	if err != nil {
		utils.ResponseJSON(w, http.StatusUnauthorized, map[string]any{"error": "Unauthorized"})
		return
	}

	if err := h.followService.AcceptFollow(payload.FollowerID, payload.FollowingID, currentUserID); err != nil {
		utils.ResponseJSON(w, http.StatusInternalServerError, map[string]any{"error": err.Error()})
		return
	}

	utils.ResponseJSON(w, http.StatusOK, map[string]any{
		"success": true,
		"message": "Follow request accepted successfully.",
	})
}

func (h *FollowHandler) DeclineFollow(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		utils.ResponseJSON(w, http.StatusMethodNotAllowed, map[string]any{"error": "Method not allowed"})
		return
	}

	var payload struct {
		FollowerID  int `json:"follower_id"`
		FollowingID int `json:"following_id"`
	}

	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		utils.ResponseJSON(w, http.StatusBadRequest, map[string]any{"error": "Invalid request body"})
		return
	}

	currentUserID, err := utils.GetCurrentUserID(r, h.sessionService)
	if err != nil {
		utils.ResponseJSON(w, http.StatusUnauthorized, map[string]any{"error": "Unauthorized"})
		return
	}

	if err := h.followService.DeclineFollow(payload.FollowerID, payload.FollowingID, currentUserID); err != nil {
		utils.ResponseJSON(w, http.StatusInternalServerError, map[string]any{"error": err.Error()})
		return
	}

	utils.ResponseJSON(w, http.StatusOK, map[string]any{
		"success": true,
		"message": "Follow request declined successfully.",
	})
}

func (h *FollowHandler) DeleteFollow(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		utils.ResponseJSON(w, http.StatusMethodNotAllowed, map[string]any{"error": "Method not allowed"})
		return
	}

	var payload struct {
		FollowerID  int `json:"follower_id"`
		FollowingID int `json:"following_id"`
	}

	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		utils.ResponseJSON(w, http.StatusBadRequest, map[string]any{"error": "Invalid request body"})
		return
	}

	currentUserID, err := utils.GetCurrentUserID(r, h.sessionService)
	if err != nil {
		utils.ResponseJSON(w, http.StatusUnauthorized, map[string]any{"error": "Unauthorized"})
		return
	}

	if err := h.followService.DeleteFollow(payload.FollowerID, payload.FollowingID, currentUserID); err != nil {
		utils.ResponseJSON(w, http.StatusInternalServerError, map[string]any{"error": err.Error()})
		return
	}

	utils.ResponseJSON(w, http.StatusOK, map[string]any{
		"success": true,
		"message": "Follow relationship deleted successfully.",
	})
}

func (h *FollowHandler) GetStatusFollow(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		utils.ResponseJSON(w, http.StatusMethodNotAllowed, map[string]any{"error": "Method not allowed"})
		return
	}

	followerIDStr := r.URL.Query().Get("follower_id")
	followingIDStr := r.URL.Query().Get("following_id")

	followerID, err := strconv.Atoi(followerIDStr)
	if err != nil {
		utils.ResponseJSON(w, http.StatusBadRequest, map[string]any{"error": "Invalid follower_id"})
		return
	}
	followingID, err := strconv.Atoi(followingIDStr)
	if err != nil {
		utils.ResponseJSON(w, http.StatusBadRequest, map[string]any{"error": "Invalid following_id"})
		return
	}

	if followerID == 0 || followingID == 0 {
		utils.ResponseJSON(w, http.StatusBadRequest, map[string]any{"error": "Follower and following IDs must be provided"})
		return
	}

	status, err := h.followService.GetStatusFollow(followerID, followingID)
	if err != nil {
		utils.ResponseJSON(w, http.StatusInternalServerError, map[string]any{"error": err.Error()})
		return
	}

	utils.ResponseJSON(w, http.StatusOK, map[string]any{
		"status":  status,
		"success": true,
	})
}

func (h *FollowHandler) GetFollowers(w http.ResponseWriter, r *http.Request) {
	fmt.Println("GetFollowers called")
	if r.Method != http.MethodGet {
		utils.ResponseJSON(w, http.StatusMethodNotAllowed, map[string]any{"error": "Method not allowed"})
		return
	}

	userIDStr := r.URL.Query().Get("user_id")
	userID, err := strconv.Atoi(userIDStr)
	if err != nil || userID <= 0 {
		utils.ResponseJSON(w, http.StatusBadRequest, map[string]any{"error": "Invalid user_id"})
		return
	}

	followers, err := h.followService.GetFollowers(userID)
	if err != nil {
		utils.ResponseJSON(w, http.StatusInternalServerError, map[string]any{"error": err.Error()})
		return
	}

	utils.ResponseJSON(w, http.StatusOK, followers)
}

func (h *FollowHandler) GetFollowing(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		utils.ResponseJSON(w, http.StatusMethodNotAllowed, map[string]any{"error": "Method not allowed"})
		return
	}

	userIDStr := r.URL.Query().Get("user_id")
	userID, err := strconv.Atoi(userIDStr)
	if err != nil || userID <= 0 {
		utils.ResponseJSON(w, http.StatusBadRequest, map[string]any{"error": "Invalid user_id"})
		return
	}

	following, err := h.followService.GetFollowing(userID)
	if err != nil {
		utils.ResponseJSON(w, http.StatusInternalServerError, map[string]any{"error": err.Error()})
		return
	}

	utils.ResponseJSON(w, http.StatusOK, following)
}