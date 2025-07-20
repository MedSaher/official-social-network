package models

// Create a user object to represent the posts model:
type Post struct {
	ID        int    `json:"id"`
	Title     string `json:"title"`         // Title added for compatibility
	Content   string `json:"content"`
	ImagePath string `json:"image_path"`    // Optional image path
	Privacy   string `json:"privacy"`       // "public", "almost_private", "private"
	CreatedAt string `json:"created_at"`    // Stored as datetime string
	UpdatedAt string `json:"updated_at"`    // Stored as datetime string
	UserId    int    `json:"user_id"`
	GroupId   *int   `json:"group_id"`      // Nullable, so we use a pointer
}

type PostUser struct {
	ID         int    `json:"id"`
	Title      string `json:"title"`
	Content    string `json:"content"`
	CreatedAt  string `json:"created_at"`
	UserId     int    `json:"user_id"`
	UserName   string `json:"user_name"`
	Categories []int  `json:"categories"` // Category IDs for post creation
	Catego     string `json:"categories_names"` // Category names as string for display
}

type Categories struct {
	ID       int    `json:"id"`
	Category string `json:"c_name"`
}
