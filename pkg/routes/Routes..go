// routes/router.go
package routes

import (
	"chat-app/pkg/handlers"
	"chat-app/pkg/middlewares"
	"chat-app/pkg/repository"
	"chat-app/pkg/services"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

// SetupRouter initializes the Gin router with all routes for the application.
func SetupRouter(db *sqlx.DB) *gin.Engine {
	router := gin.Default()

	// Initialize the repository
	userRepo := repository.NewUserRepository(db)
	chatRepo := repository.NewMessageRepository(db)
	// Initialize the services
	userService := services.NewUserService(userRepo)
	chatService := services.NewChatService(chatRepo)
	// Initialize the handlers
	userHandler := handlers.NewUserHandler(userService)
	chatHandler := handlers.NewChatHandler(chatService)
	// Public routes
	router.POST("/login", userHandler.Login)
	router.POST("/register", userHandler.Register)

	// Setup protected routes using the AuthMiddleware
	authMiddleware := middlewares.AuthMiddleware(os.Getenv("JWT_SECRET"))
	protected := router.Group("/")
	protected.Use(authMiddleware)
	{
		// Add protected routes here
		// e.g., protected.GET("/user/profile", userHandler.GetProfile)
		protected.POST("/users/:senderID/messages/:recipientID/send", chatHandler.SendMessage)
		protected.GET("/users/:senderID/messages/:recipientID", chatHandler.GetConversation)
	}

	return router
}
