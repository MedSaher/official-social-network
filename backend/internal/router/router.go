package router

import (
	"net/http"
	"strings"

	"social_network/internal/services"
)

type Router struct {
	Routes        map[string]http.HandlerFunc
	usersSessions services.SessionsServicesLayer
}

func NewRouter(session *services.SessionService) *Router {
	return &Router{
		Routes:        make(map[string]http.HandlerFunc),
		usersSessions: session,
	}
}

func (router *Router) AddRoute(method string, path string, handler http.HandlerFunc) {
	route := strings.ToLower(method + ":" + path)
	router.Routes[route] = handler
}

func (router *Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	frontEndPaths := map[string]bool{
		"/register": true,
		"/login":    true,
	}

	// CORS headers
origin := r.Header.Get("Origin")

// Only allow trusted origins — never blindly echo in production
w.Header().Set("Access-Control-Allow-Origin", origin)
w.Header().Set("Access-Control-Allow-Credentials", "true") // 🔥 REQUIRED for credentials
w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}

	// Check session
	isValid := false
	session, err := r.Cookie("session_token")
	if err == nil && session != nil {
		isValid = router.usersSessions.IsValidSession(session.Value)
	}

	// If the user is logged in and tries to go to login or register -> redirect to home
	if frontEndPaths[r.URL.Path] && r.Method == "GET" {
		if isValid {
			http.Redirect(w, r, "/", http.StatusFound)
		} else {
			http.ServeFile(w, r, "../frontend/index.html")
		}
		return
	}

	// If user is not logged in and tries to go to home
	if r.Method == "GET" && (r.URL.Path == "/" || r.URL.Path == "/index.html") {
		if !isValid {
			http.Redirect(w, r, "/login", http.StatusFound)
			return
		}
		http.ServeFile(w, r, "../frontend/index.html")
		return
	}

	// Serve static files
	if r.Method == "GET" {
		if strings.HasSuffix(r.URL.Path, ".css") || strings.HasSuffix(r.URL.Path, ".js") || strings.HasSuffix(r.URL.Path, ".png") {
			http.ServeFile(w, r, "../frontend"+r.URL.Path)
			return
		}
		// http.ServeFile(w, r, "../front/index.html")
		// return
	}

	// Handle registered routes
	route := strings.ToLower(r.Method + ":" + r.URL.Path)
	if handler, ok := router.Routes[route]; ok {
		handler(w, r)
		return
	}

	// Not found
	http.ServeFile(w, r, "../frontend/index.html")
}
