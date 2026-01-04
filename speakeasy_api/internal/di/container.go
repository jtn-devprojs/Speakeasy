package di

import (
	"database/sql"

	_ "modernc.org/sqlite"

	"github.com/speakeasy/speakeasy-api/internal/config"
	"github.com/speakeasy/speakeasy-api/internal/controllers"
	"github.com/speakeasy/speakeasy-api/internal/repositories"
	"github.com/speakeasy/speakeasy-api/internal/services"
)

type Container struct {
	DB                *sql.DB
	ConfigLoader      config.ConfigLoader
	UserRepo          repositories.IUserRepository
	SessionRepo       repositories.ISessionRepository
	SessionUserRepo   repositories.ISessionUserRepository
	MessageRepo       repositories.IMessageRepository
	AuthService       services.IAuthService
	SessionService    services.ISessionService
	SessionController *controllers.SessionController
}

func NewContainer(db *sql.DB, dbType string) *Container {
	configLoader := &config.DefaultConfigLoader{}
	userRepo := repositories.NewUserRepository(db)
	sessionRepo := repositories.NewSessionRepository(db)

	// Select session locker based on database type
	var locker repositories.ISessionLocker
	switch dbType {
	case "postgres":
		locker = &repositories.PostgresSessionLocker{}
	case "sqlite":
		locker = &repositories.SqliteSessionLocker{}
	default:
		locker = &repositories.SqliteSessionLocker{}
	}

	sessionUserRepo := repositories.NewSessionUserRepository(db, locker)
	messageRepo := repositories.NewMessageRepository(db)

	authService := services.NewAuthService(userRepo)
	sessionService := services.NewSessionService(sessionRepo, sessionUserRepo)

	sessionController := controllers.NewSessionController(sessionService)

	return &Container{
		DB:                db,
		ConfigLoader:      configLoader,
		UserRepo:          userRepo,
		SessionRepo:       sessionRepo,
		SessionUserRepo:   sessionUserRepo,
		MessageRepo:       messageRepo,
		AuthService:       authService,
		SessionService:    sessionService,
		SessionController: sessionController,
	}
}
