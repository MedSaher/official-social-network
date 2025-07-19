package main

import (
	"log"
	"net/http"
	"social-app-backend/db"
	"social-app-backend/handlers"

	gorillaHandlers "github.com/gorilla/handlers" // alias to avoid conflict
	"github.com/gorilla/mux"
)

func main() {
	db.Init()

	r := mux.NewRouter()
	r.HandleFunc("/register", handlers.Register).Methods("POST")
	r.HandleFunc("/login", handlers.Login).Methods("POST")

	log.Println("Server running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", gorillaHandlers.CORS(
		gorillaHandlers.AllowedOrigins([]string{"http://localhost:3000"}),
		gorillaHandlers.AllowedMethods([]string{"GET", "POST", "OPTIONS"}),
		gorillaHandlers.AllowedHeaders([]string{"Content-Type"}),
	)(r)))
}
