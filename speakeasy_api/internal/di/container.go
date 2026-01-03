package di

import (
	"database/sql"

	_ "modernc.org/sqlite"

	"github.com/speakeasy/speakeasy-api/internal/controllers"
	"github.com/speakeasy/speakeasy-api/internal/repositories"
	"github.com/speakeasy/speakeasy-api/internal/services"
)

type Container struct {
	DB                *sql.DB
	UserRepo          repositories.IUserRepository
	SessionRepo       repositories.ISessionRepository
	MessageRepo       repositories.IMessageRepository
	UserService       services.IUserService
	AuthService       services.IAuthService
	LocationService   services.ILocationService
	UserController    *controllers.UserController
	AuthController    *controllers.AuthController
	SessionController *controllers.SessionController
}

func NewContainer(db *sql.DB) *Container {
	userRepo := repositories.NewUserRepository(db)
	sessionRepo := repositories.NewSessionRepository(db)
	messageRepo := repositories.NewMessageRepository(db)

	userService := services.NewUserService(userRepo)
	authService := services.NewAuthService(userRepo)
	locationService := services.NewLocationService(sessionRepo)

	userController := controllers.NewUserController(userService)
	authController := controllers.NewAuthController(authService)
	sessionController := controllers.NewSessionController(locationService)

	return &Container{
		DB:                db,
		UserRepo:          userRepo,
		SessionRepo:       sessionRepo,
		MessageRepo:       messageRepo,
		UserService:       userService,
		AuthService:       authService,
		LocationService:   locationService,
		UserController:    userController,
		AuthController:    authController,
		SessionController: sessionController,
	}
}
