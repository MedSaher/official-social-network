package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"social_network/internal/adapters/http/utils"
	"social_network/internal/domain/ports/service"
	"strconv"
	"strings"
)

type ProfileHandler struct {
	profileService service.UserService
	sessionService service.SessionService
}

func NewProfileHandler(profileSvc service.UserService, sessionSvc service.SessionService) *ProfileHandler {
	return &ProfileHandler{
		profileService: profileSvc,
		sessionService: sessionSvc,
	}
}

func (p *ProfileHandler) Profile(w http.ResponseWriter, r *http.Request) {
	fmt.Println("actived")
	path := r.URL.Path
	parts := strings.Split(path, "/")
	if len(parts) < 4 {
		http.Error(w, "Invalid URL", http.StatusBadRequest)
		return
	}
	id := parts[3]
	fmt.Println(id)
	userId, err := strconv.Atoi(id)
	if err != nil {
		utils.ResponseJSON(w, http.StatusNotFound, map[string]any{"error": "Page Not Found"})
		return
	}

	// Fake data (simulate DB fetch)
	profile, err := p.profileService.GetProfile(userId)
	if err != nil {
		utils.ResponseJSON(w, http.StatusNotFound, map[string]any{"error": "Page Not Found"})
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(profile)
}
