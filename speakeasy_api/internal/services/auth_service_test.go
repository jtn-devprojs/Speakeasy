package services

import (
	"database/sql"
	"testing"

	"github.com/speakeasy/speakeasy-api/internal/repositories"
	_ "modernc.org/sqlite"
)

func createTestAuthService(t *testing.T) *AuthService {
	db, err := sql.Open("sqlite", ":memory:")
	if err != nil {
		t.Fatalf("Failed to create test database: %v", err)
	}

	schema := `
	CREATE TABLE IF NOT EXISTS users (
		id TEXT PRIMARY KEY,
		username TEXT NOT NULL UNIQUE,
		email TEXT,
		avatar_url TEXT,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	);
	`

	_, err = db.Exec(schema)
	if err != nil {
		t.Fatalf("Failed to create schema: %v", err)
	}

	userRepo := repositories.NewUserRepository(db)
	return NewAuthService(userRepo)
}

func TestNewAuthService(t *testing.T) {
	authService := createTestAuthService(t)

	if authService == nil {
		t.Fatal("Expected non-nil AuthService")
	}
}

func TestAuthService_Login(t *testing.T) {
	authService := createTestAuthService(t)

	_, err := authService.Login("testuser", "password")
	if err != ErrNotImplemented {
		t.Fatalf("Expected ErrNotImplemented, got %v", err)
	}
}

func TestAuthService_Logout(t *testing.T) {
	authService := createTestAuthService(t)

	err := authService.Logout("token")
	if err != ErrNotImplemented {
		t.Fatalf("Expected ErrNotImplemented, got %v", err)
	}
}

func TestAuthService_Register(t *testing.T) {
	authService := createTestAuthService(t)

	_, err := authService.Register("testuser", "test@example.com", "password")
	if err != ErrNotImplemented {
		t.Fatalf("Expected ErrNotImplemented, got %v", err)
	}
}

func TestAuthService_ValidateToken(t *testing.T) {
	authService := createTestAuthService(t)

	_, err := authService.ValidateToken("token")
	if err != ErrNotImplemented {
		t.Fatalf("Expected ErrNotImplemented, got %v", err)
	}
}

func TestAuthService_RefreshToken(t *testing.T) {
	authService := createTestAuthService(t)

	_, err := authService.RefreshToken("token")
	if err != ErrNotImplemented {
		t.Fatalf("Expected ErrNotImplemented, got %v", err)
	}
}
