package handlers

import (
	"net/http"
	"social_network/internal/handlers/utils"
	"social_network/internal/services"
)

// Create the handler for the websocket:
type WebSocketHandler struct {
	socketService services.WebsocketSeviceLayer
	sessionServ   services.SessionsServicesLayer
}

// Create a new instance of the websocket handler:
func NewWebSocketHandler(socketServ *services.WebSocketService, sessionServ *services.SessionService) *WebSocketHandler {
	return &WebSocketHandler{
		socketService: socketServ,
		sessionServ:   sessionServ,
	}
}

// Request structure for marking messages as read
type MarkAsReadRequest struct {
	SenderID int `json:"sender_id"`
}

func (soc *WebSocketHandler) SocketHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		utils.ResponseJSON(w, http.StatusMethodNotAllowed, map[string]any{"message": "invalid method"})
		return
	}
	if r.Header.Get("Upgrade") != "websocket" {
		utils.ResponseJSON(w, http.StatusBadRequest, map[string]any{"message": "Expected WebSocket upgrade"})
		return
	}
	if err := soc.socketService.CreateNewWebSocket(w, r); err != nil {
		utils.ResponseJSON(w, http.StatusInternalServerError, map[string]any{"message": "failed to create websocket"})
		return
	}
}

func (soc *WebSocketHandler) GetUsers(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("session_token")
	if err != nil {
		utils.ResponseJSON(w, http.StatusUnauthorized, map[string]any{"message": "No session token found"})
		return
	}

	userID, err := soc.sessionServ.GetUserIdFromSession(cookie.Value)
	if err != nil {
		utils.ResponseJSON(w, http.StatusUnauthorized, map[string]any{"message": "Invalid session"})
		return
	}

	// Extract offset and limit from query parameters, with fallback default values:

	offset, limit := utils.ParseLimitOffset(r)

	// Call the service with offset and limit parameters:
	users, err := soc.socketService.GetAllUsersWithStatus(userID, offset, limit)
	if err != nil {
		utils.ResponseJSON(w, http.StatusInternalServerError, map[string]any{"error": err.Error()})
		return
	}
	utils.ResponseJSON(w, http.StatusOK, users)
}
