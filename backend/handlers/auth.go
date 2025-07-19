package handlers

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"social-app-backend/db"
	"social-app-backend/models"

	"golang.org/x/crypto/bcrypt"
)

func Register(w http.ResponseWriter, r *http.Request) {
	var user models.User
	json.NewDecoder(r.Body).Decode(&user)

	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		http.Error(w, "Server error", http.StatusInternalServerError)
		return
	}

	stmt, err := db.DB.Prepare("INSERT INTO users(email, password) VALUES(?, ?)")
	if err != nil {
		http.Error(w, "Database error", http.StatusInternalServerError)
		return
	}
	_, err = stmt.Exec(user.Email, string(hash))
	if err != nil {
		http.Error(w, "Email already exists", http.StatusConflict)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode("User registered")
}

func Login(w http.ResponseWriter, r *http.Request) {
	var input models.User
	json.NewDecoder(r.Body).Decode(&input)

	var stored models.User
	err := db.DB.QueryRow("SELECT id, password FROM users WHERE email = ?", input.Email).Scan(&stored.Id, &stored.Password)
	if err == sql.ErrNoRows {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(stored.Password), []byte(input.Password))
	if err != nil {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}

	// Could return JWT token or session ID here
	json.NewEncoder(w).Encode(map[string]interface{}{"message": "Login successful", "userId": stored.Id})
}
