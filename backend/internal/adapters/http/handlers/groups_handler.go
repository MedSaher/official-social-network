package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"social_network/internal/adapters/http/utils"
	"social_network/internal/domain/models"
	"social_network/internal/domain/ports/service"
)

type GroupHandler struct {
	groupService service.GroupService
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

