package repositories

import (
	"database/sql"
	"testing"

	_ "modernc.org/sqlite"
)

func createTestUserRepo(t *testing.T) *UserRepository {
	db, err := sql.Open("sqlite3", ":memory:")
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

	return NewUserRepository(db)
}

func TestNewUserRepository(t *testing.T) {
	repo := createTestUserRepo(t)

	if repo == nil {
		t.Fatal("Expected non-nil UserRepository")
	}
}

func TestUserRepository_CreateUser(t *testing.T) {
	repo := createTestUserRepo(t)

	err := repo.CreateUser("user123", "testuser", nil, nil)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
}

func TestUserRepository_GetUser(t *testing.T) {
	repo := createTestUserRepo(t)

	repo.CreateUser("user123", "testuser", nil, nil)

	user, err := repo.GetUser("user123")
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	if user == nil {
		t.Fatal("Expected user to be found")
	}

	if user.ID != "user123" {
		t.Fatalf("Expected ID 'user123', got %s", user.ID)
	}
}

func TestUserRepository_GetUserNotFound(t *testing.T) {
	repo := createTestUserRepo(t)

	user, err := repo.GetUser("non-existent")
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	if user != nil {
		t.Fatal("Expected nil for non-existent user")
	}
}

func TestUserRepository_GetUserByUsername(t *testing.T) {
	repo := createTestUserRepo(t)

	repo.CreateUser("user123", "testuser", nil, nil)

	user, err := repo.GetUserByUsername("testuser")
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	if user == nil {
		t.Fatal("Expected user to be found")
	}

	if user.Username != "testuser" {
		t.Fatalf("Expected username 'testuser', got %s", user.Username)
	}
}

func TestUserRepository_UpdateUser(t *testing.T) {
	repo := createTestUserRepo(t)

	repo.CreateUser("user123", "testuser", nil, nil)

	err := repo.UpdateUser("user123", "updateduser", nil, nil)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	user, _ := repo.GetUser("user123")
	if user.Username != "updateduser" {
		t.Fatalf("Expected username 'updateduser', got %s", user.Username)
	}
}

func TestUserRepository_DeleteUser(t *testing.T) {
	repo := createTestUserRepo(t)

	repo.CreateUser("user123", "testuser", nil, nil)

	err := repo.DeleteUser("user123")
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	user, _ := repo.GetUser("user123")
	if user != nil {
		t.Fatal("Expected user to be deleted")
	}
}
