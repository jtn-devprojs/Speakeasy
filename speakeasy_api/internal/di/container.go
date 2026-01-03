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
	SessionUserRepo   repositories.ISessionUserRepository
	MessageRepo       repositories.IMessageRepository
	UserService       services.IUserService
	AuthService       services.IAuthService
	SessionService    services.ISessionService
	UserController    *controllers.UserController
	AuthController    *controllers.AuthController
	SessionController *controllers.SessionController
}

func NewContainer(db *sql.DB) *Container {
	userRepo := repositories.NewUserRepository(db)
	sessionRepo := repositories.NewSessionRepository(db)
	sessionUserRepo := repositories.NewSessionUserRepository(db)
	messageRepo := repositories.NewMessageRepository(db)

	userService := services.NewUserService(userRepo)
	authService := services.NewAuthService(userRepo)
	sessionService := services.NewSessionService(sessionRepo, sessionUserRepo)

	userController := controllers.NewUserController(userService)
	authController := controllers.NewAuthController(authService)
	sessionController := controllers.NewSessionController(sessionService)

	return &Container{
		DB:                db,
		UserRepo:          userRepo,
		SessionRepo:       sessionRepo,
		SessionUserRepo:   sessionUserRepo,
		MessageRepo:       messageRepo,
		UserService:       userService,
		AuthService:       authService,
		SessionService:    sessionService,
		UserController:    userController,
		AuthController:    authController,
		SessionController: sessionController,
	}
}
