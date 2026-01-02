package repositories

import (
	"database/sql"
	"testing"

	_ "modernc.org/sqlite"
)

func createTestSessionRepo(t *testing.T) *SessionRepository {
	db, err := sql.Open("sqlite", ":memory:")
	if err != nil {
		t.Fatalf("Failed to create test database: %v", err)
	}

	schema := `
	CREATE TABLE IF NOT EXISTS sessions (
		location TEXT PRIMARY KEY,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		status TEXT DEFAULT 'active'
	);
	`

	_, err = db.Exec(schema)
	if err != nil {
		t.Fatalf("Failed to create schema: %v", err)
	}

	return NewSessionRepository(db)
}

func TestNewSessionRepository(t *testing.T) {
	repo := createTestSessionRepo(t)

	if repo == nil {
		t.Fatal("Expected non-nil SessionRepository")
	}
}

func TestSessionRepository_CreateSession(t *testing.T) {
	repo := createTestSessionRepo(t)

	err := repo.CreateSession("Test Location", "active")
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
}

func TestSessionRepository_GetSession(t *testing.T) {
	repo := createTestSessionRepo(t)

	repo.CreateSession("Test Location", "active")

	session, err := repo.GetSession("Test Location")
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	if session == nil {
		t.Fatal("Expected session to be found")
	}

	if session.Location != "Test Location" {
		t.Fatalf("Expected location 'Test Location', got %s", session.Location)
	}
}

func TestSessionRepository_GetSessionNotFound(t *testing.T) {
	repo := createTestSessionRepo(t)

	session, err := repo.GetSession("Non-existent")
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	if session != nil {
		t.Fatal("Expected nil for non-existent session")
	}
}

func TestSessionRepository_UpdateSessionStatus(t *testing.T) {
	repo := createTestSessionRepo(t)

	repo.CreateSession("Test Location", "active")

	err := repo.UpdateSessionStatus("Test Location", "archived")
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	session, _ := repo.GetSession("Test Location")
	if session.Status != "archived" {
		t.Fatalf("Expected status 'archived', got %s", session.Status)
	}
}

func TestSessionRepository_DeleteSession(t *testing.T) {
	repo := createTestSessionRepo(t)

	repo.CreateSession("Test Location", "active")

	err := repo.DeleteSession("Test Location")
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	session, _ := repo.GetSession("Test Location")
	if session != nil {
		t.Fatal("Expected session to be deleted")
	}
}
