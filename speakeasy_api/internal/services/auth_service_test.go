package services

import (
	"testing"
)

func TestNewAuthService(t *testing.T) {
	userService := NewUserService()
	authService := NewAuthService(userService)

	if authService == nil {
		t.Fatal("Expected non-nil AuthService")
	}
}

func TestAuthService_Login(t *testing.T) {
	userService := NewUserService()
	authService := NewAuthService(userService)

	_, err := authService.Login("testuser", "password")
	if err != ErrNotImplemented {
		t.Fatalf("Expected ErrNotImplemented, got %v", err)
	}
}

func TestAuthService_Logout(t *testing.T) {
	userService := NewUserService()
	authService := NewAuthService(userService)

	err := authService.Logout("token")
	if err != ErrNotImplemented {
		t.Fatalf("Expected ErrNotImplemented, got %v", err)
	}
}

func TestAuthService_Register(t *testing.T) {
	userService := NewUserService()
	authService := NewAuthService(userService)

	_, err := authService.Register("testuser", "test@example.com", "password")
	if err != ErrNotImplemented {
		t.Fatalf("Expected ErrNotImplemented, got %v", err)
	}
}

func TestAuthService_ValidateToken(t *testing.T) {
	userService := NewUserService()
	authService := NewAuthService(userService)

	_, err := authService.ValidateToken("token")
	if err != ErrNotImplemented {
		t.Fatalf("Expected ErrNotImplemented, got %v", err)
	}
}

func TestAuthService_RefreshToken(t *testing.T) {
	userService := NewUserService()
	authService := NewAuthService(userService)

	_, err := authService.RefreshToken("token")
	if err != ErrNotImplemented {
		t.Fatalf("Expected ErrNotImplemented, got %v", err)
	}
}
