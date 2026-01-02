package di

import (
	"github.com/speakeasy/speakeasy-api/internal/controllers"
	"github.com/speakeasy/speakeasy-api/internal/services"
)

// Container holds all application dependencies
type Container struct {
	UserService        *services.UserService
	AuthService        *services.AuthService
	LocationService    *services.LocationService
	UserController     *controllers.UserController
	AuthController     *controllers.AuthController
	LocationController *controllers.LocationController
}

// NewContainer initializes and returns a new dependency injection container
func NewContainer() *Container {
	// Initialize services
	userService := services.NewUserService()
	authService := services.NewAuthService(userService)
	locationService := services.NewLocationService()

	// Initialize controllers
	userController := controllers.NewUserController(userService)
	authController := controllers.NewAuthController(authService)
	locationController := controllers.NewLocationController(locationService)

	return &Container{
		UserService:        userService,
		AuthService:        authService,
		LocationService:    locationService,
		UserController:     userController,
		AuthController:     authController,
		LocationController: locationController,
	}
}
