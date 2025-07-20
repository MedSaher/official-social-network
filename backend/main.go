package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"social_network/db/sqlite"
	"social_network/internal/handlers"
	"social_network/internal/repositories"
	"social_network/internal/router"
	"social_network/internal/services"
)

var (
	databaseConnection *sql.DB
	mainError          error
)

func init() {
	// Database setup (path and migration)
	dbPath := "./data/social_network.db"
	migrationsPath := "./db/migrations/sqlite"

	// Connect and run migrations
	databaseConnection, mainError = sqlite.ConnectAndMigrate(dbPath, migrationsPath)
	if mainError != nil {
		log.Fatalf("Failed to connect and migrate DB: %v", mainError)
	}
	log.Println("Database connection successful")
}

func main() {
	// Ensure that no database connection errors occurred during init()
	if mainError != nil {
		log.Fatalf("Database connection failed: %v", mainError)
		return
	}
	defer databaseConnection.Close() // Ensure DB is closed at the end of the program
	fmt.Println("Connected successfully to the database")

	// Chat Broker setup
	chatBroker := services.NewChatBroker()
	go chatBroker.RunChatBroker()

	// Repositories initialization
	usersRepository := repositories.NewUsersRepository(databaseConnection)
	sessionRepository := repositories.NewSessionsRepository(databaseConnection)
	messageRepository := repositories.NewMessageRepository(databaseConnection)

	// Services initialization
	usersServices := services.NewUsersServices(usersRepository)
	sessionService := services.NewSessionsServices(usersRepository, sessionRepository)
	webSocketService := services.NewWebSocketService(chatBroker, messageRepository, sessionRepository, usersRepository)
	// messagesService := services.NewMessageService(messageRepository, sessionRepository)

	// Handlers initialization
	usersHandlers := handlers.NewUsersHandlers(chatBroker, usersServices, sessionService)
	webSocketHandler := handlers.NewWebSocketHandler(webSocketService, sessionService)
	// messagesHandler := handlers.NewMessagesHandler(messagesService, sessionService)

	// Router initialization and routes setup
	mainRouter := router.NewRouter(sessionService)

	// Authentication routes
	mainRouter.AddRoute("POST", "/api/login", usersHandlers.UsersLoginHandler)
	mainRouter.AddRoute("POST", "/api/signup", usersHandlers.UsersRegistrationHandler)
	// mainRouter.AddRoute("POST", "/api/logout", usersHandlers.UsersLogoutHandler)
	mainRouter.AddRoute("GET", "/api/check-session", usersHandlers.CheckSessionHandler)


	// websocket and chat routes:
	mainRouter.AddRoute("GET", "/api/ws", webSocketHandler.SocketHandler)

	// Print message indicating the server is listening
	fmt.Println("Listening on port: http://localhost:8080/")

	// Start the server and handle incoming requests
	mainError = http.ListenAndServe(":8080", mainRouter)
	if mainError != nil {
		log.Fatalf("Error starting the server: %v", mainError)
	}
}

/* package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"social_network/db/sqlite"
	"social_network/internal/handlers"
	"social_network/internal/repositories"
	"social_network/internal/router"
	"social_network/internal/services"
)

var (
	databaseConnection *sql.DB
	mainError          error
)

func init() {
	dbPath := "./data/social_network.db"
	migrationsPath := "./db/migrations/sqlite"

	databaseConnection, err := sqlite.ConnectAndMigrate(dbPath, migrationsPath)
	if err != nil {
		log.Fatalf("Failed to connect and migrate DB: %v", err)
	}
	defer databaseConnection.Close()

	// Now you can use `db` to query your SQLite database.

	log.Println("Backend app started")
}

func main() {
	if mainError != nil {
		return
	}
	defer databaseConnection.Close()
	fmt.Println("Connected successfully to database")

	// craete Chat Broker :
	chatBroker := services.NewChatBroker()
	go chatBroker.RunChatBroker()

	// Initialize repositories:
	usersRepository := repositories.NewUsersRepository(databaseConnection)
	sessionRepository := repositories.NewSessionsRepository(databaseConnection)
	// postsRepository := repositories.NewPostsRepository(databaseConnection)
	// commentsRepository := repositories.NewCommentsRepository(databaseConnection)
	// messageRepository := repositories.NewMessageRepository(databaseConnection)

	// Initialize services:
	usersServices := services.NewUsersServices(usersRepository)
	sessionService := services.NewSessionsServices(usersRepository, sessionRepository)
	// postsServices := services.NewPostService(postsRepository, sessionRepository)
	// commentsService := services.NewCommentsServices(commentsRepository, sessionRepository)
	// webSocketService := services.NewWebSocketService(chatBroker, messageRepository, sessionRepository, usersRepository)
	// messagesService := services.NewMessageService(messageRepository, sessionRepository)

	// Initialize handlers:
	usersHandlers := handlers.NewUsersHandlers(chatBroker, usersServices, sessionService)
	// postsHandlers := handlers.NewPostsHandles(postsServices)
	// commentsHandlers := handlers.NewCommentsHandler(commentsService)
	// webSocketHandler := handlers.NewWebSocketHandler(webSocketService, sessionService)
	// messagesHandler := handlers.NewMessagesHandler(messagesService, sessionService)
	// Setup router and routes:
	mainRouter := router.NewRouter(sessionService)


// Authentication routes
mainRouter.AddRoute("POST", "/api/login", usersHandlers.UsersLoginHandler)
mainRouter.AddRoute("POST", "/api/signup", usersHandlers.UsersRegistrationHandler)
mainRouter.AddRoute("POST", "/api/logout", usersHandlers.UsersLogoutHandler)
mainRouter.AddRoute("GET", "/api/check-session", usersHandlers.UsersCheckSessionHandler)

// // User routes
// mainRouter.AddRoute("GET", "/api/user/", usersHandlers.GetUserHandler)
// mainRouter.AddRoute("GET", "/api/users", usersHandlers.GetUsersHandler)
// mainRouter.AddRoute("GET", "/api/images", usersHandlers.GetImageHandler)
// mainRouter.AddRoute("POST", "/api/user/update", usersHandlers.UpdateUserHandler)
// mainRouter.AddRoute("GET", "/api/top-engaged-users", usersHandlers.GetTopEngagedUsersHandler)
// mainRouter.AddRoute("GET", "/api/user/posts", postsHandlers.GetUserPostsHandler)

// // Post routes
// mainRouter.AddRoute("GET", "/api/posts", postsHandlers.GetAllPostsHandler)
// mainRouter.AddRoute("POST", "/api/post", postsHandlers.CreatePostHandler)
// mainRouter.AddRoute("POST", "/api/post/update", postsHandlers.UpdatePostHandler)
// mainRouter.AddRoute("POST", "/api/post/delete", postsHandlers.DeletePostHandler)
// mainRouter.AddRoute("GET", "/api/post/single", postsHandlers.GetSinglePostHandler)

// // Group posts
// mainRouter.AddRoute("POST", "/api/group/post", postsHandlers.CreateGroupPostHandler)
// mainRouter.AddRoute("GET", "/api/group/posts", postsHandlers.GetGroupPostsHandler)

// // Comment routes
// mainRouter.AddRoute("GET", "/api/comments", commentsHandlers.GetCommentsHandler)
// mainRouter.AddRoute("POST", "/api/comment", commentsHandlers.AddCommentHandler)
// mainRouter.AddRoute("POST", "/api/comment/update", commentsHandlers.UpdateCommentHandler)
// // mainRouter.AddRoute("POST", "/api/comment/delete", commentsHandlers.DeleteCommentHandler) // still commented

// // Reaction routes
// mainRouter.AddRoute("POST", "/api/react", middleware.AuthMiddleware(reactionsHandlers.ReactHandler))
// mainRouter.AddRoute("GET", "/api/reactions", middleware.AuthMiddleware(reactionsHandlers.GetAvailableReactionsHandler))

// // Group routes
// mainRouter.AddRoute("GET", "/api/groups", middleware.AuthMiddleware(groupsHandlers.GetGroupsHandler))
// mainRouter.AddRoute("GET", "/api/group", middleware.AuthMiddleware(groupsHandlers.GetGroupHandler))
// mainRouter.AddRoute("GET", "/api/group/details", middleware.AuthMiddleware(groupsHandlers.GetGroupDetailsHandler))
// mainRouter.AddRoute("POST", "/api/group-requests", groupsHandlers.GroupRequestHandler)
// mainRouter.AddRoute("POST", "/api/group-joinRequests", groupsHandlers.GroupJoinRequestHandler)
// mainRouter.AddRoute("POST", "/api/group/create", groupsHandlers.CreateGroupHandler)
// mainRouter.AddRoute("POST", "/api/group/update", middleware.AuthMiddleware(groupsHandlers.UpdateGroupHandler))
// mainRouter.AddRoute("POST", "/api/group/delete", middleware.AuthMiddleware(groupsHandlers.DeleteGroupHandler))

// // Follow routes
// mainRouter.AddRoute("POST", "/api/follow", followHandlers.InitFollowHandler)
// mainRouter.AddRoute("GET", "/api/following/", followHandlers.FollowingHandler)
// mainRouter.AddRoute("GET", "/api/followers/", followHandlers.FollowersHandler)
// mainRouter.AddRoute("POST", "/api/follow-requests", followHandlers.FollowRequestHandler)
// mainRouter.AddRoute("GET", "/api/followers", followHandlers.GetFollowersHandler)

// // Event routes
// mainRouter.AddRoute("GET", "/api/events", eventsHandlers.GetEventsHandler)
// mainRouter.AddRoute("POST", "/api/event", eventsHandlers.CreateEventHandler)
// mainRouter.AddRoute("POST", "/api/event/respond", eventsHandlers.RespondToEventHandler)
// mainRouter.AddRoute("GET", "/api/event/responses", eventsHandlers.GetEventResponsesHandler)

// // Notification routes
// mainRouter.AddRoute("GET", "/api/notifications", notificationsHandlers.GetAllNotificationsHandler)
// mainRouter.AddRoute("GET", "/api/new-notifications", notificationsHandlers.GetNewNotificationsHandler)
// mainRouter.AddRoute("POST", "/api/notification/read", notificationsHandlers.MarkNotificationReadHandler)
// mainRouter.AddRoute("GET", "/api/notifications/unread-count", notificationsHandlers.NotificationCountUnreadHandler)

// // Chat routes
// mainRouter.AddRoute("GET", "/api/ws", chatHandlers.InitWebSocketConnectionHandler)
// mainRouter.AddRoute("GET", "/api/chat", chatHandlers.GetChatHandler)
// mainRouter.AddRoute("GET", "/api/chat-group", chatHandlers.GetGroupChatHandler)
// mainRouter.AddRoute("GET", "/api/chats", chatHandlers.GetAllChatsHandler)
// mainRouter.AddRoute("GET", "/api/chat/newusers", chatHandlers.GetNewChatUsersHandler)
// mainRouter.AddRoute("POST", "/api/chat/send", chatHandlers.SendMessageHandler)
// mainRouter.AddRoute("POST", "/api/chat/mark-read", chatHandlers.MarkMessageAsReadHandler)

// // Group invitation routes
// mainRouter.AddRoute("POST", "/api/group/invite", middleware.AuthMiddleware(invitationHandlers.InviteUsersHandler))
// mainRouter.AddRoute("POST", "/api/group/cancel-invite", middleware.AuthMiddleware(invitationHandlers.CancelInvitationHandler))
// mainRouter.AddRoute("GET", "/api/group/invite-list", middleware.AuthMiddleware(invitationHandlers.GetGroupInvitationListHandler))
// mainRouter.AddRoute("POST", "/api/group/invite/accept", invitationHandlers.AcceptInvitationHandler)
// mainRouter.AddRoute("POST", "/api/group/invite/reject", invitationHandlers.RejectInvitationHandler)
// mainRouter.AddRoute("POST", "/api/group/members/remove", invitationHandlers.RemoveMembersHandler)


	// fmt.Println("Routes registered:", mainRouter.Routes)
	fmt.Println("Listening on port: http://localhost:8080/")

	mainError = http.ListenAndServe(":8080", mainRouter)
	if mainError != nil {
		return
	}
}
*/
