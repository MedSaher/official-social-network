package router

import (
	"net/http"
	"strings"

	"social_network/internal/domain/ports/service"
)

type Router struct {
	routes      map[string]http.HandlerFunc
	sessionServ service.SessionService
	staticFiles map[string]http.Handler
}

// NewRouter crée un nouveau routeur avec session + static frontend
func NewRouter(session service.SessionService) *Router {
	return &Router{
		routes:      make(map[string]http.HandlerFunc),
		sessionServ: session,
		staticFiles: map[string]http.Handler{
			"/frontend/": http.StripPrefix("/frontend/", http.FileServer(http.Dir("./frontend/out"))),
			"/avatars/":  http.StripPrefix("/avatars/", http.FileServer(http.Dir("./avatars"))),
		},
	}
}

// AddRoute permet d'ajouter une route (ex: POST:/api/register)
func (r *Router) AddRoute(method string, path string, handler http.HandlerFunc) {
	key := strings.ToLower(method + ":" + path)
	r.routes[key] = handler
}

// ServeHTTP gère les requêtes entrantes
func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	// CORS
	origin := req.Header.Get("Origin")
	w.Header().Set("Access-Control-Allow-Origin", origin)
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

	if req.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}

	// Serve static files by prefix
	for prefix, handler := range r.staticFiles {
		if strings.HasPrefix(req.URL.Path, prefix) {
			handler.ServeHTTP(w, req)
			return
		}
	}

	// ✅ NEW: Serve /uploads/posts/* static files
	if strings.HasPrefix(req.URL.Path, "/uploads/posts/") {
		fs := http.StripPrefix("/uploads/posts/", http.FileServer(http.Dir("./uploads/posts")))
		fs.ServeHTTP(w, req)
		return
	}

	// Handle API routes
	if strings.HasPrefix(req.URL.Path, "/api/") {
		key := strings.ToLower(req.Method + ":" + req.URL.Path)
		if handler, ok := r.routes[key]; ok {
			handler(w, req)
			return
		}
		http.Error(w, "API endpoint not found", http.StatusNotFound)
		return
	}

	// Default: frontend static files (Next.js build)
	if indexHandler, ok := r.staticFiles["/frontend/"]; ok {
		indexHandler.ServeHTTP(w, req)
		return
	}

	http.NotFound(w, req)
}
