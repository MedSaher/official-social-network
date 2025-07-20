package handlers

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"social_network/internal/handlers/utils"
	"social_network/internal/services"
)

// PostsHandlers handles all post-related HTTP logic
type PostsHandlers struct {
	postsServ services.PostsServiceLayer
}

// NewPostsHandlers creates a new instance of PostsHandlers
func NewPostsHandlers(postSer *services.PostsService) *PostsHandlers {
	return &PostsHandlers{
		postsServ: postSer,
	}
}

// CreatePostsHandler handles creating a new post
func (postHand *PostsHandlers) CreatePostsHandler(w http.ResponseWriter, r *http.Request) {
	// Only allow POST requests
	if r.Method != http.MethodPost {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	// Max upload size (e.g., 10MB)
	r.ParseMultipartForm(10 << 20)

	// Get form values
	title := r.FormValue("title")
	content := r.FormValue("content")
	privacy := r.FormValue("privacy")

	// Get file from "image" input
	file, fileHeader, err := r.FormFile("image")
	var imagePath string
	if err == nil {
		defer file.Close()

		// Create a folder if it doesn't exist
		os.MkdirAll("uploads", os.ModePerm)

		// Save the file with a unique name
		filename := fmt.Sprintf("%d_%s", time.Now().Unix(), fileHeader.Filename)
		imagePath = filepath.Join("uploads", filename)

		dst, err := os.Create(imagePath)
		if err != nil {
			http.Error(w, "Unable to save image", http.StatusInternalServerError)
			return
		}
		defer dst.Close()

		_, err = io.Copy(dst, file)
		if err != nil {
			http.Error(w, "Failed to save image", http.StatusInternalServerError)
			return
		}
	}

	// Optional: Save to database
	// For now, just log the data
	fmt.Println("New post:")
	fmt.Println("Title:", title)
	fmt.Println("Content:", content)
	fmt.Println("Privacy:", privacy)
	fmt.Println("Image Path:", imagePath)

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Post created successfully"))
}

// if r.Method != http.MethodPost {
// 	utils.ResponseJSON(w, http.StatusMethodNotAllowed, map[string]any{"message": "Invalid method"})
// 	return
// }

// var post models.PostUser
// if err := json.NewDecoder(r.Body).Decode(&post); err != nil {
// 	utils.ResponseJSON(w, http.StatusBadRequest, map[string]any{"message": "Invalid request body"})
// 	return
// }

// session, err := r.Cookie("session_token")
// if err != nil || session == nil {
// 	utils.ResponseJSON(w, http.StatusUnauthorized, map[string]any{"message": "Missing or invalid session token"})
// 	return
// }

// if err := postHand.postsServ.CreatePost(&post, session.Value); err != nil {
// 	utils.ResponseJSON(w, http.StatusInternalServerError, map[string]any{"message": "Failed to create post"})
// 	return
// }

// utils.ResponseJSON(w, http.StatusCreated, map[string]string{"message": "Post created successfully"})

// GetAllPostsHandler handles fetching all posts with pagination
func (postHand *PostsHandlers) GetAllPostsHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		utils.ResponseJSON(w, http.StatusMethodNotAllowed, map[string]any{"message": "Method not allowed"})
		return
	}

	offset, limit := utils.ParseLimitOffset(r)
	posts, err := postHand.postsServ.GetAllPostsService(offset, limit)
	if err != nil {
		utils.ResponseJSON(w, http.StatusInternalServerError, map[string]any{"message": "Failed to fetch posts"})
		return
	}

	utils.ResponseJSON(w, http.StatusOK, posts)
}

// GetAllCategoriesHandler handles fetching all available post categories
func (postHand *PostsHandlers) GetAllCategoriesHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		utils.ResponseJSON(w, http.StatusMethodNotAllowed, map[string]any{"message": "Method not allowed"})
		return
	}

	categories, err := postHand.postsServ.GetAllCategoriesService()
	if err != nil {
		utils.ResponseJSON(w, http.StatusInternalServerError, map[string]any{"message": "Failed to fetch categories"})
		return
	}

	utils.ResponseJSON(w, http.StatusOK, categories)
}
