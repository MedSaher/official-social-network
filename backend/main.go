package main

import (
	"log"
	"net/http"

	"social_network/internal/adapters/http/router"
	"social_network/internal/adapters/http/handlers"
	"social_network/internal/infrastructure/db/sqlite"
	"social_network/internal/infrastructure/db/repository"
	"social_network/internal/application/service"
)

func main() {
	dbPath := "./social_network/social_network.db"
	migrationsPath := "./internal/infrastructure/db/migrations"

	db, err := sqlite.ConnectDB(dbPath)
	if err != nil {
		log.Fatal("DB connection error:", err)
	}

	if err := sqlite.RunMigrations(db, migrationsPath); err != nil {
		log.Fatal("Migration error:", err)
	}

	// Repositories
	userRepo := repository.NewUserRepository(db)
	sessionRepo := repository.NewSessionRepository(db)

	// Services
	userService := service.NewUserService(userRepo)
	sessionService := service.NewSessionService(sessionRepo)

	// Handlers
	userHandler := handlers.NewUserHandler(userService, sessionService)

	// Router
	r := router.NewRouter()

	// Register routes
	r.POST("/register", userHandler.Register)
	r.POST("/login", userHandler.Login)
	r.POST("/logout", userHandler.Logout)

	// Start server
	log.Println("🚀 Server running on http://localhost:8080")
	http.ListenAndServe(":8080", r)
}
