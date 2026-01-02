package services

import (
	"database/sql"
	"testing"

	"github.com/speakeasy/speakeasy-api/internal/repositories"
	_ "modernc.org/sqlite"
)

func createTestUserRepo(t *testing.T) *repositories.UserRepository {
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

	return repositories.NewUserRepository(db)
}

func TestNewUserService(t *testing.T) {
	repo := createTestUserRepo(t)
	service := NewUserService(repo)

	if service == nil {
		t.Fatal("Expected non-nil UserService")
	}
}

func TestUserService_GetUserByID(t *testing.T) {
	repo := createTestUserRepo(t)
	service := NewUserService(repo)

	result, err := service.GetUserByID("user123")
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	if result != nil {
		t.Fatal("Expected nil result for non-existent user")
	}
}

func TestUserService_DeleteUser(t *testing.T) {
	repo := createTestUserRepo(t)
	service := NewUserService(repo)

	err := service.DeleteUser("user123")
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
}

func TestUserService_GetUserPreferences(t *testing.T) {
	mockRepo := &repositories.MockUserRepository{}
	service := NewUserService(mockRepo)

	_, err := service.GetUserPreferences("user123")
	if err != ErrNotImplemented {
		t.Fatalf("Expected ErrNotImplemented, got %v", err)
	}
}

func TestUserService_UpdateUserPreferences(t *testing.T) {
	mockRepo := &repositories.MockUserRepository{}
	service := NewUserService(mockRepo)

	err := service.UpdateUserPreferences("user123", map[string]interface{}{})
	if err != ErrNotImplemented {
		t.Fatalf("Expected ErrNotImplemented, got %v", err)
	}
}
