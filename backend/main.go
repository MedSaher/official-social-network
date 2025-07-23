package main

import (
	"log"
	"net/http"

	"social_network/internal/adapters/db"
	"social_network/internal/adapters/http/handlers"
	"social_network/internal/adapters/router"
	"social_network/internal/application/services"
	"social_network/internal/infrastructure/db/sqlite"
)

func main() {
	dbPath := "./social_network/social_network.db"
	migrationsPath := "./internal/infrastructure/db/migrations"

	sqliteDB, err := sqlite.ConnectDB(dbPath)
	if err != nil {
		log.Fatal("DB connection error:", err)
	}

	if err := sqlite.RunMigrations(sqliteDB, migrationsPath); err != nil {
		log.Fatal("Migration error:", err)
	}

	// Repositories
	userRepo := db.NewUserRepository(sqliteDB)
	sessionRepo := db.NewSessionRepository(sqliteDB)
	followRepo := db.NewFollowRepository(sqliteDB)

	// Services
	userService := services.NewUserService(userRepo)
	sessionService := services.NewSessionService(userRepo, sessionRepo)
	followService := services.NewFollowService(followRepo,userRepo)

	// Handlers
	userHandler := handlers.NewUserHandler(userService, sessionService)
	followHandler := handlers.NewFollowHandler(followService, sessionService)

	// Router
	r := router.NewRouter(sessionService)

	// auth routes
	r.AddRoute("POST", "/api/register", userHandler.Register)
	r.AddRoute("POST", "/api/login", userHandler.Login)
	r.AddRoute("POST", "/api/logout", userHandler.Logout)

	// Follow routes
	r.AddRoute("POST", "/api/follow", followHandler.CreateFollow)
	r.AddRoute("POST", "/api/follow/accept", followHandler.AcceptFollow)
	r.AddRoute("POST", "/api/follow/decline", followHandler.DeclineFollow)
	
	// Start server
	log.Println("🚀 Server running on http://localhost:8080")
	http.ListenAndServe(":8080", r)
}
