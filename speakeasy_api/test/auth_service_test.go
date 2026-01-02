package main

import (
	"testing"

	"github.com/speakeasy/speakeasy-api/internal/services"
)

func TestNewAuthService(t *testing.T) {
	userService := services.NewUserService()
	authService := services.NewAuthService(userService)

	if authService == nil {
		t.Fatal("Expected non-nil AuthService")
	}
}

func TestAuthServiceLogin(t *testing.T) {
	userService := services.NewUserService()
	authService := services.NewAuthService(userService)

	_, err := authService.Login("testuser", "password")
	if err != services.ErrNotImplemented {
		t.Fatalf("Expected ErrNotImplemented, got %v", err)
	}
}

func TestAuthServiceLogout(t *testing.T) {
	userService := services.NewUserService()
	authService := services.NewAuthService(userService)

	err := authService.Logout("token")
	if err != services.ErrNotImplemented {
		t.Fatalf("Expected ErrNotImplemented, got %v", err)
	}
}

func TestAuthServiceRegister(t *testing.T) {
	userService := services.NewUserService()
	authService := services.NewAuthService(userService)

	_, err := authService.Register("testuser", "test@example.com", "password")
	if err != services.ErrNotImplemented {
		t.Fatalf("Expected ErrNotImplemented, got %v", err)
	}
}

func TestAuthServiceValidateToken(t *testing.T) {
	userService := services.NewUserService()
	authService := services.NewAuthService(userService)

	_, err := authService.ValidateToken("token")
	if err != services.ErrNotImplemented {
		t.Fatalf("Expected ErrNotImplemented, got %v", err)
	}
}

func TestAuthServiceRefreshToken(t *testing.T) {
	userService := services.NewUserService()
	authService := services.NewAuthService(userService)

	_, err := authService.RefreshToken("token")
	if err != services.ErrNotImplemented {
		t.Fatalf("Expected ErrNotImplemented, got %v", err)
	}
}
