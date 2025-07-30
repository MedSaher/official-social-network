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
		fmt.Println("error fetching groups:", err)
		utils.ResponseJSON(w, http.StatusInternalServerError, map[string]any{"error": "Failed to fetch groups"})
		return
	}

	utils.ResponseJSON(w, http.StatusOK, groups)
}

// inside handlers/group_handler.go

func (h *GroupHandler) DynamicRoutes(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.URL.Path)
	if strings.HasSuffix(r.URL.Path, "/join") {
		h.JoinGroup(w, r)
		return
	} else if strings.HasSuffix(r.URL.Path, "/pending_requests") {
		h.FetchPendingRequests(w, r)
		return
	} else if strings.HasSuffix(r.URL.Path, "/member_role") {
		h.GetMemberRole(w, r)
		return
	} else if strings.HasSuffix(r.URL.Path, "/posts") {
		h.GetGroupPosts(w, r)
		return
	} else if strings.HasSuffix(r.URL.Path, "/events") {
		h.GetGroupEvents(w, r)  
		return
	} else if strings.HasSuffix(r.URL.Path, "/group_info") {
		h.GetGroupInfo(w, r)
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

// internal/adapters/http/handlers/groups.go

func (h *GroupHandler) RespondToJoinRequest(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		utils.ResponseJSON(w, http.StatusMethodNotAllowed, map[string]any{"error": "Method Not Allowed"})
		return
	}

	var req struct {
		RequestID int  `json:"request_id"`
		Accept    bool `json:"accept"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		utils.ResponseJSON(w, http.StatusBadRequest, map[string]any{"error": "Invalid JSON"})
		return
	}

	userID, err := utils.GetCurrentUserID(r, h.sessionService)
	if err != nil {
		utils.ResponseJSON(w, http.StatusUnauthorized, map[string]any{"error": "Unauthorized"})
		return
	}

	err = h.groupService.RespondToJoinRequest(r.Context(), req.RequestID, userID, req.Accept)
	if err != nil {
		utils.ResponseJSON(w, http.StatusForbidden, map[string]any{"error": err.Error()})
		return
	}

	utils.ResponseJSON(w, http.StatusOK, map[string]string{"message": "Request handled successfully"})
}

func (h *GroupHandler) GetMemberRole(w http.ResponseWriter, r *http.Request) {
	userID, err := utils.GetCurrentUserID(r, h.sessionService)
	if err != nil {
		utils.ResponseJSON(w, http.StatusUnauthorized, map[string]any{"error": "unauthorized"})
		return
	}

	parts := strings.Split(r.URL.Path, "/")
	if len(parts) < 5 {
		utils.ResponseJSON(w, http.StatusBadRequest, map[string]any{"error": "Invalid group ID"})
		return
	}
	groupIDStr := parts[3]

	groupID, err := strconv.Atoi(groupIDStr)
	if err != nil {
		utils.ResponseJSON(w, http.StatusBadRequest, map[string]any{"error": "invalid group id"})
		return
	}

	role, err := h.groupService.GetUserRole(r.Context(), groupID, userID)
	if err != nil {
		utils.ResponseJSON(w, http.StatusInternalServerError, map[string]any{"error": "failed to fetch role"})
		return
	}

	utils.ResponseJSON(w, http.StatusOK, map[string]string{"role": role})
}

func (h *GroupHandler) GetGroupPosts(w http.ResponseWriter, r *http.Request) {
	_, err := utils.GetCurrentUserID(r, h.sessionService)
	if err != nil {
		utils.ResponseJSON(w, http.StatusUnauthorized, map[string]any{"error": "unauthorized"})
		return
	}
	parts := strings.Split(r.URL.Path, "/")
	if len(parts) < 5 {
		utils.ResponseJSON(w, http.StatusBadRequest, map[string]any{"error": "Invalid group ID"})
		return
	}
	groupIDStr := parts[3]

	groupID, err := strconv.Atoi(groupIDStr)
	if err != nil {
		utils.ResponseJSON(w, http.StatusBadRequest, map[string]any{"error": "invalid group id"})
		return
	}
	posts, err := h.groupService.GetGroupPosts(r.Context(), groupID)
	if err != nil {
		utils.ResponseJSON(w, http.StatusInternalServerError, map[string]any{"error": "failed to fetch posts"})
		return
	}

	utils.ResponseJSON(w, http.StatusOK, posts)
}

func (h *GroupHandler) GetGroupEvents(w http.ResponseWriter, r *http.Request) {
	_, err := utils.GetCurrentUserID(r, h.sessionService)
	if err != nil {
		utils.ResponseJSON(w, http.StatusUnauthorized, map[string]any{"error": "unauthorized"})
		return
	}
	parts := strings.Split(r.URL.Path, "/")
	if len(parts) < 5 {
		utils.ResponseJSON(w, http.StatusBadRequest, map[string]any{"error": "Invalid group ID"})
		return
	}
	groupIDStr := parts[3]

	groupID, err := strconv.Atoi(groupIDStr)
	if err != nil {
		utils.ResponseJSON(w, http.StatusBadRequest, map[string]any{"error": "invalid group id"})
		return
	}
	events, err := h.groupService.GetGroupEvents(r.Context(), groupID)
	if err != nil {
		utils.ResponseJSON(w, http.StatusInternalServerError, map[string]any{"error": "failed to fetch events"})
		return
	}
	utils.ResponseJSON(w, http.StatusOK, events)
}

func (h *GroupHandler) GetGroupInfo(w http.ResponseWriter, r *http.Request) {
	// Extract group ID from URL
	parts := strings.Split(r.URL.Path, "/")
	if len(parts) < 5 {
		utils.ResponseJSON(w, http.StatusBadRequest, map[string]any{"error": "Invalid group ID"})
		return
	}
	groupID, err := strconv.Atoi(parts[3])
	if err != nil {
		utils.ResponseJSON(w, http.StatusBadRequest, map[string]any{"error": "Invalid group ID"})
		return
	}

	info, err := h.groupService.GetGroupInfo(r.Context(), groupID)
	if err != nil {
		utils.ResponseJSON(w, http.StatusInternalServerError, map[string]any{"error": "Failed to fetch group info"})
		return
	}

	utils.ResponseJSON(w, http.StatusOK, info)
}
