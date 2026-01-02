package di

import (
	"github.com/speakeasy/speakeasy-api/internal/handlers"
	"github.com/speakeasy/speakeasy-api/internal/services"
)

// Container holds all application dependencies
type Container struct {
	UserService *services.UserService
	AuthService *services.AuthService
	UserHandler *handlers.UserHandler
	AuthHandler *handlers.AuthHandler
}

// NewContainer initializes and returns a new dependency injection container
func NewContainer() *Container {
	// Initialize services
	userService := services.NewUserService()
	authService := services.NewAuthService(userService)

	// Initialize handlers
	userHandler := handlers.NewUserHandler(userService)
	authHandler := handlers.NewAuthHandler(authService)

	return &Container{
		UserService: userService,
		AuthService: authService,
		UserHandler: userHandler,
		AuthHandler: authHandler,
	}
}
