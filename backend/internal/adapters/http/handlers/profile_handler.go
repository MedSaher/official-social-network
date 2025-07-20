package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"
)

type Profile struct {
	ID          string `json:"id"`
	FirstName   string `json:"firstName"`
	LastName    string `json:"lastName"`
	Nickname    string `json:"nickname"`
	Avatar      string `json:"avatar"`
	AboutMe     string `json:"aboutMe"`
	IsPrivate   bool   `json:"isPrivate"`
	IsOwnProfile bool  `json:"isOwnProfile"`
	IsFollowing bool   `json:"isFollowing"`
}

func (p *Profile) Profile(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path // Example: /api/profile/123
	parts := strings.Split(path, "/")
	if len(parts) < 4 {
		http.Error(w, "Invalid URL", http.StatusBadRequest)
		return
	}
	id := parts[3] // Extract ID

	// Fake data (simulate DB fetch)
	profile := Profile{
		ID:          id,
		FirstName:   "Saher",
		LastName:    "Mohamed",
		Nickname:    "saher_dev",
		Avatar:      "https://i.pravatar.cc/150?img=5",
		AboutMe:     "Go programmer and cybersecurity enthusiast",
		IsPrivate:   false,
		IsOwnProfile: false,
		IsFollowing: true,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(profile)
}

func main() {
	http.HandleFunc("/api/profile/", getProfileHandler)

	log.Println("Server running at http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
