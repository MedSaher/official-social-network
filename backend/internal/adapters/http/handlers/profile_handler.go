package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"social_network/internal/domain/models"
	"social_network/internal/domain/ports/service"
	"strings"
)

type ProfileHandler struct {
	profileService service.ProfileService
	sessionService service.SessionService
}



func NewProfileHandler(profileSvc service.ProfileService, sessionSvc service.SessionService) *ProfileHandler {
	return &ProfileHandler{
		profileService: profileSvc,
		sessionService: sessionSvc,
	}
}

func (p *ProfileHandler) Profile(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path // Example: /api/profile/123
	parts := strings.Split(path, "/")
	if len(parts) < 4 {
		http.Error(w, "Invalid URL", http.StatusBadRequest)
		return
	}
	id := parts[3] // Extract ID

	fmt.Println(id)

	// Fake data (simulate DB fetch)
	profile := models.Profile{
		ID:           id,
		FirstName:    "Saher",
		LastName:     "Mohamed",
		Nickname:     "saher_dev",
		Avatar:       "https://i.pravatar.cc/150?img=5",
		AboutMe:      "Go programmer and cybersecurity enthusiast",
		IsPrivate:    false,
		IsOwnProfile: false,
		IsFollowing:  true,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(profile)
}
