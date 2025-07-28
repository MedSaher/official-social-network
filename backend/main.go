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
	sessionRepo := db.NewSessionRepository(sqliteDB)
	followRepo := db.NewFollowRepository(sqliteDB)
	userRepo := db.NewUserRepository(sqliteDB)
	postRepo := db.NewPostRepository(sqliteDB)
	groupsRepo := db.NewGroupRepository(sqliteDB)

	// Services
	userService := services.NewUserService(userRepo, followRepo,postRepo)
	sessionService := services.NewSessionService(userRepo, sessionRepo)
	followService := services.NewFollowService(followRepo, userRepo)
	postService := services.NewPostService(postRepo)
	groupsService := services.NewGroupService(groupsRepo)

	// Handlers
	userHandler := handlers.NewUserHandler(userService, sessionService)
	followHandler := handlers.NewFollowHandler(followService, sessionService)
	postHandler := handlers.NewPostHandler(postService, sessionService)
	groupsHandler := handlers.NewGroupHandler(groupsService, sessionService)

	// Router
	r := router.NewRouter(sessionService)

	// auth routes
	r.AddRoute("POST", "/api/register", userHandler.Register)
	r.AddRoute("POST", "/api/login", userHandler.Login)
	r.AddRoute("POST", "/api/logout", userHandler.Logout)
	r.AddRoute("GET", "/api/check-session", userHandler.CheckSession)

	//profile
	r.AddRoute("GET", "/api/profile", userHandler.GetFullProfile)
	r.AddRoute("POST", "/api/profile/privacy", userHandler.ChangePrivacyStatus)
	r.AddRoute("GET", "/api/search_users", userHandler.SearchUsers)
	r.AddRoute("GET", "/api/user/by_username", userHandler.GetUserProfileByUsername)

	// Follow routes
	r.AddRoute("POST", "/api/follow", followHandler.CreateFollow)
	r.AddRoute("POST", "/api/follow/accept", followHandler.AcceptFollow)
	r.AddRoute("POST", "/api/follow/decline", followHandler.DeclineFollow)
	r.AddRoute("DELETE", "/api/follow/delete", followHandler.DeleteFollow)
	r.AddRoute("GET", "/api/follow/status", followHandler.GetStatusFollow)
	r.AddRoute("GET", "/api/follow/followers", followHandler.GetFollowers)
	r.AddRoute("GET", "/api/follow/following", followHandler.GetFollowing)

	// posts routes
	r.AddRoute("POST", "/api/posts/create_comment", postHandler.CreateComment)
	r.AddRoute("POST", "/api/posts/create_post", postHandler.CreatePost)
	r.AddRoute("GET", "/api/posts/fetch_posts", postHandler.GetPosts)
	r.AddRoute("GET", "/api/posts/fetch_comments", postHandler.FetchComments)

	// groupes routes
	r.AddRoute("POST", "/api/groups/create_group", groupsHandler.CreateGroup)
	r.AddRoute("GET", "/api/groups/fetch_groups", groupsHandler.FetchGroups)
	r.AddPrefixRoute("POST", "/api/groups/", groupsHandler.DynamicRoutes)
	r.AddRoute("POST", "/api/groups/join_request/respond", groupsHandler.RespondToJoinRequest)
	// r.AddPrefixRoute("GET", "/api/groups/", groupsHandler.DynamicRoutes)

	// Start server
	log.Println("🚀 Server running on http://localhost:8080")
	err = http.ListenAndServe(":8080", r)
	if err != nil {
		log.Fatal("Server error:", err)
	}
}
