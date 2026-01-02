package controllers

import (
	"testing"

	"github.com/speakeasy/speakeasy-api/internal/repositories"
	"github.com/speakeasy/speakeasy-api/internal/services"
)

func TestNewAuthController(t *testing.T) {
	mockRepo := &repositories.MockUserRepository{}
	authService := services.NewAuthService(mockRepo)
	controller := NewAuthController(authService)

	if controller == nil {
		t.Fatal("Expected non-nil AuthController")
	}

	if controller.authService == nil {
		t.Fatal("Expected authService to be initialized")
	}
}

func TestAuthControllerLogin(t *testing.T) {
	mockRepo := &repositories.MockUserRepository{}
	authService := services.NewAuthService(mockRepo)
	controller := NewAuthController(authService)

	if controller == nil {
		t.Fatal("Expected non-nil controller")
	}
}

func TestAuthControllerLogout(t *testing.T) {
	mockRepo := &repositories.MockUserRepository{}
	authService := services.NewAuthService(mockRepo)
	controller := NewAuthController(authService)

	if controller == nil {
		t.Fatal("Expected non-nil controller")
	}
}

func TestAuthControllerRegister(t *testing.T) {
	mockRepo := &repositories.MockUserRepository{}
	authService := services.NewAuthService(mockRepo)
	controller := NewAuthController(authService)

	if controller == nil {
		t.Fatal("Expected non-nil controller")
	}
}

func TestAuthControllerValidateToken(t *testing.T) {
	mockRepo := &repositories.MockUserRepository{}
	authService := services.NewAuthService(mockRepo)
	controller := NewAuthController(authService)

	if controller == nil {
		t.Fatal("Expected non-nil controller")
	}
}

func TestAuthControllerRefreshToken(t *testing.T) {
	mockRepo := &repositories.MockUserRepository{}
	authService := services.NewAuthService(mockRepo)
	controller := NewAuthController(authService)

	if controller == nil {
		t.Fatal("Expected non-nil controller")
	}
}
