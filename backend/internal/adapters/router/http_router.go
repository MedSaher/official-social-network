package router

import (
	"net/http"
	"strings"

	"social_network/internal/domain/ports/service"
)

type Router struct {
	routes      map[string]http.HandlerFunc
	sessionServ service.SessionService
	staticFiles http.Handler
}

// NewRouter crée un nouveau routeur avec session + static frontend
func NewRouter(session service.SessionService) *Router {
	return &Router{
		routes:      make(map[string]http.HandlerFunc),
		sessionServ: session,
		staticFiles: http.FileServer(http.Dir("./frontend/out")), // <-- frontend exporté ici
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
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	if req.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}

	// API: /api/... → dispatch vers les handlers
	if strings.HasPrefix(req.URL.Path, "/api/") {
		key := strings.ToLower(req.Method + ":" + req.URL.Path)
		if handler, ok := r.routes[key]; ok {
			handler(w, req)
			return
		}
		http.Error(w, "API endpoint not found", http.StatusNotFound)
		return
	}

	// Autres → fichiers statiques (Next.js build)
	r.staticFiles.ServeHTTP(w, req)
}
