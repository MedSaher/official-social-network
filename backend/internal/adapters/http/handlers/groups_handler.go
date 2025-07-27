package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"social_network/internal/adapters/http/utils"
	"social_network/internal/domain/models"
	"social_network/internal/domain/ports/service"
)

type GroupHandler struct {
	groupService   service.GroupService
	sessionService service.SessionService
}

func NewGroupHandler(s service.GroupService, ses service.SessionService) *GroupHandler {
	return &GroupHandler{groupService: s, sessionService: ses}
}

func (h *GroupHandler) CreateGroup(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		utils.ResponseJSON(w, http.StatusMethodNotAllowed, map[string]any{"error": "Method Not Allowed"})
		return
	}

	var group models.Group
	if err := json.NewDecoder(r.Body).Decode(&group); err != nil {
		utils.ResponseJSON(w, http.StatusBadRequest, map[string]any{"error": "Invalid JSON"})
		return
	}

	userID, err := utils.GetCurrentUserID(r, h.sessionService)
	fmt.Println("user : ", userID)
	if err != nil {
		utils.ResponseJSON(w, http.StatusUnauthorized, map[string]any{"error": "Unauthorized"})
		return
	}

	if group.Title == "" {
		utils.ResponseJSON(w, http.StatusBadRequest, map[string]any{"error": "Title is required"})
		return
	}

	// ✅ assign the user ID from session
	group.CreatorID = userID

	if err := h.groupService.CreateGroup(r.Context(), &group); err != nil {
		utils.ResponseJSON(w, http.StatusInternalServerError, map[string]any{"error": "Failed to create group"})
		return
	}

	utils.ResponseJSON(w, http.StatusCreated, group)
}

func (h *GroupHandler) FetchGroups(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		utils.ResponseJSON(w, http.StatusMethodNotAllowed, map[string]any{"error": "Method Not Allowed"})
		return
	}

	userID, err := utils.GetCurrentUserID(r, h.sessionService)
	if err != nil {
		utils.ResponseJSON(w, http.StatusUnauthorized, map[string]any{"error": "Unauthorized"})
		return
	}

	groups, err := h.groupService.GetGroupsForUser(r.Context(), userID)
	if err != nil {
		utils.ResponseJSON(w, http.StatusInternalServerError, map[string]any{"error": "Failed to fetch groups"})
		return
	}

	utils.ResponseJSON(w, http.StatusOK, groups)
}

// inside handlers/group_handler.go

func (h *GroupHandler) DynamicRoutes(w http.ResponseWriter, r *http.Request) {
	if strings.HasSuffix(r.URL.Path, "/join") {
		h.JoinGroup(w, r)
		return
		} else if strings.HasSuffix(r.URL.Path, "/pending_requests") {
		h.FetchPendingRequests(w, r)
		return
	}
}


func (h *GroupHandler) JoinGroup(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		utils.ResponseJSON(w, http.StatusMethodNotAllowed, map[string]any{"error": "Method Not Allowed"})
		return
	}

	userID, err := utils.GetCurrentUserID(r, h.sessionService)
	if err != nil {
		utils.ResponseJSON(w, http.StatusUnauthorized, map[string]any{"error": "Unauthorized"})
		return
	}

	// Extract group ID from URL path
	path := r.URL.Path
	parts := strings.Split(path, "/")
	if len(parts) < 5 {
		utils.ResponseJSON(w, http.StatusBadRequest, map[string]any{"error": "Invalid group ID in path"})
		return
	}
	groupIDStr := parts[3]
	groupID, err := strconv.Atoi(groupIDStr)
	if err != nil {
		fmt.Println("error : ", err)
		utils.ResponseJSON(w, http.StatusBadRequest, map[string]any{"error": "Invalid group ID"})
		return
	}

	// Call service
	err = h.groupService.RequestToJoinGroup(r.Context(), groupID, userID)
	if err != nil {
		utils.ResponseJSON(w, http.StatusConflict, map[string]any{"error": err.Error()})
		return
	}

	utils.ResponseJSON(w, http.StatusCreated, map[string]string{
		"message": "Join request sent",
	})
}

func (h *GroupHandler) FetchPendingRequests(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		utils.ResponseJSON(w, http.StatusMethodNotAllowed, map[string]any{"error": "Method Not Allowed"})
		return
	}

	userID, err := utils.GetCurrentUserID(r, h.sessionService)
	if err != nil {
		utils.ResponseJSON(w, http.StatusUnauthorized, map[string]any{"error": "Unauthorized"})
		return
	}

	// Extract group ID from URL path
	parts := strings.Split(r.URL.Path, "/")
	if len(parts) < 5 {
		utils.ResponseJSON(w, http.StatusBadRequest, map[string]any{"error": "Invalid group ID"})
		return
	}
	groupID, err := strconv.Atoi(parts[3])
	if err != nil {
		fmt.Println(err)
		utils.ResponseJSON(w, http.StatusBadRequest, map[string]any{"error": "Invalid group ID"})
		return
	}

	// Verify user is creator of the group
	isCreator, err := h.groupService.IsCreator(r.Context(), groupID, userID)
	if err != nil {
		fmt.Println(err)

		utils.ResponseJSON(w, http.StatusInternalServerError, map[string]any{"error": "Internal error"})
		return
	}
	if !isCreator {
		fmt.Println(err)

		utils.ResponseJSON(w, http.StatusForbidden, map[string]any{"error": "Access denied"})
		return
	}

	requests, err := h.groupService.GetPendingRequests(r.Context(), groupID)
	if err != nil {
		fmt.Println(err)

		utils.ResponseJSON(w, http.StatusInternalServerError, map[string]any{"error": "Failed to fetch requests"})
		return
	}

	utils.ResponseJSON(w, http.StatusOK, requests)
}
