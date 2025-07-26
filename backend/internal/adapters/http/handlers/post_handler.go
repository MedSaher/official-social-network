package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"social_network/internal/adapters/http/utils"
	"social_network/internal/domain/models"
	"social_network/internal/domain/ports/service"
	"strconv"
	"strings"
	"time"
)

type PostHandler struct {
	postService    service.PostService
	sessionService service.SessionService
}

func NewPostHandler(postSvc service.PostService, sessionSvc service.SessionService) *PostHandler {
	return &PostHandler{
		postService:    postSvc,
		sessionService: sessionSvc,
	}
}

func (p *PostHandler) CreatePost(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		utils.ResponseJSON(w, http.StatusMethodNotAllowed, map[string]any{"error": "Method Not Allowed"})
		return
	}
	
	// Validate & parse user_id (required)
	// Get token from "Token" header (you can change the key name if needed)
	token, err := r.Cookie("session_token")
	if token.Value == "" || err != nil {
		utils.ResponseJSON(w, http.StatusUnauthorized, map[string]any{"error": "Missing session token in headers"})
		return
	}

	userID, err := p.sessionService.GetUserIdFromSession(token.Value)
	if err != nil {
		utils.ResponseJSON(w, http.StatusUnauthorized, map[string]any{"error": "Invalid or expired session token"})
		return
	}

	// Parse multipart form, limit max upload size (e.g. 10MB)
	err = r.ParseMultipartForm(10 << 20) // 10MB
	if err != nil {
		utils.ResponseJSON(w, http.StatusBadRequest, map[string]any{"error": "Failed to parse form data"})
		return
	}

	// Parse optional group_id
	var groupID *int = nil
	groupIDStr := r.FormValue("group_id")
	if groupIDStr != "" {
		gid, err := strconv.Atoi(groupIDStr)
		if err != nil {
			utils.ResponseJSON(w, http.StatusBadRequest, map[string]any{"error": "Invalid group_id"})
			return
		}
		groupID = &gid
	}

	// Validate content (required)
	content := strings.TrimSpace(r.FormValue("content"))
	if content == "" {
		utils.ResponseJSON(w, http.StatusBadRequest, map[string]any{"error": "content is required"})
		return
	}

	// Validate privacy with allowed values
	privacy := r.FormValue("privacy")
	switch privacy {
	case "public", "almost_private", "private":
		// ok
	default:
		privacy = "public" // default fallback
	}

	// Handle optional image upload
	var savedImagePath string = ""

	file, header, err := r.FormFile("image")
	if err == nil {
		defer file.Close()

		// Validate file extension (basic)
		ext := strings.ToLower(filepath.Ext(header.Filename))
		if ext != ".jpg" && ext != ".jpeg" && ext != ".png" && ext != ".gif" {
			utils.ResponseJSON(w, http.StatusBadRequest, map[string]any{"error": "Invalid image format"})
			return
		}

		// Create uploads directory if not exists
		uploadDir := "./uploads/posts"
		if err := os.MkdirAll(uploadDir, os.ModePerm); err != nil {
			utils.ResponseJSON(w, http.StatusInternalServerError, map[string]any{"error": "Failed to create upload directory"})
			return
		}

		// Save file with unique name (timestamp + original filename)
		fileName := fmt.Sprintf("%d_%s", time.Now().UnixNano(), header.Filename)
		filePath := filepath.Join(uploadDir, fileName)

		dst, err := os.Create(filePath)
		if err != nil {
			utils.ResponseJSON(w, http.StatusInternalServerError, map[string]any{"error": "Failed to save image"})
			return
		}
		defer dst.Close()

		_, err = io.Copy(dst, file)
		if err != nil {
			utils.ResponseJSON(w, http.StatusInternalServerError, map[string]any{"error": "Failed to save image"})
			return
		}

		savedImagePath = filePath // or relative path as per your app
	} else {
		fmt.Println(err)
	}
	// Call service to create post
	post, err := p.postService.CreatePost(r.Context(), userID, groupID, content, privacy, savedImagePath)
	if err != nil {
		utils.ResponseJSON(w, http.StatusInternalServerError, map[string]any{"error": err.Error()})
		return
	}
	// Return created post as JSON
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(post)
}

func (h *PostHandler) GetPosts(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodGet {
        utils.ResponseJSON(w, http.StatusMethodNotAllowed, map[string]any{
            "error": "Method not allowed",
        })
        return
    }

    posts, err := h.postService.GetAllPosts(r.Context())
    if err != nil {
        utils.ResponseJSON(w, http.StatusInternalServerError, map[string]any{
            "error": "Failed to fetch posts",
        })
        return
    }

    utils.ResponseJSON(w, http.StatusOK, posts)
}

func (h *PostHandler) FetchComments(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodGet {
        utils.ResponseJSON(w, http.StatusMethodNotAllowed, map[string]any{"error": "Method Not Allowed"})
        return
    }

    postIDStr := r.URL.Query().Get("post_id")
    if postIDStr == "" {
        utils.ResponseJSON(w, http.StatusBadRequest, map[string]any{"error": "post_id is required"})
        return
    }

    postID, err := strconv.Atoi(postIDStr)
    if err != nil {
        utils.ResponseJSON(w, http.StatusBadRequest, map[string]any{"error": "invalid post_id"})
        return
    }

    comments, err := h.postService.GetCommentsByPostID(r.Context(), postID)
    if err != nil {
        utils.ResponseJSON(w, http.StatusInternalServerError, map[string]any{"error": "failed to fetch comments"})
        return
    }

    utils.ResponseJSON(w, http.StatusOK, comments)
}

func (h *PostHandler) CreateComment(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodPost {
        utils.ResponseJSON(w, http.StatusMethodNotAllowed, map[string]any{"error": "Method Not Allowed"})
        return
    }

    var c models.Comment
    if err := json.NewDecoder(r.Body).Decode(&c); err != nil {
        utils.ResponseJSON(w, http.StatusBadRequest, map[string]any{"error": "Invalid JSON"})
        return
    }

    if c.PostID == 0 || c.UserID == 0 || len(c.Content) == 0 {
        utils.ResponseJSON(w, http.StatusBadRequest, map[string]any{"error": "post_id, user_id, and content are required"})
        return
    }

    err := h.postService.CreateComment(r.Context(), &c)
    if err != nil {
        utils.ResponseJSON(w, http.StatusInternalServerError, map[string]any{"error": "Failed to create comment"})
        return
    }

    utils.ResponseJSON(w, http.StatusCreated, c)
}
