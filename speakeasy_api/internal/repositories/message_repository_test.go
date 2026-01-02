package repositories

import (
	"database/sql"
	"testing"

	_ "modernc.org/sqlite"
)

func createTestMessageRepo(t *testing.T) *MessageRepository {
	db, err := sql.Open("sqlite", ":memory:")
	if err != nil {
		t.Fatalf("Failed to create test database: %v", err)
	}

	schema := `
	CREATE TABLE IF NOT EXISTS messages (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		session_location TEXT NOT NULL,
		user_id TEXT NOT NULL,
		content TEXT NOT NULL,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		edited_at TIMESTAMP
	);
	`

	_, err = db.Exec(schema)
	if err != nil {
		t.Fatalf("Failed to create schema: %v", err)
	}

	return NewMessageRepository(db)
}

func TestNewMessageRepository(t *testing.T) {
	repo := createTestMessageRepo(t)

	if repo == nil {
		t.Fatal("Expected non-nil MessageRepository")
	}
}

func TestMessageRepository_CreateMessage(t *testing.T) {
	repo := createTestMessageRepo(t)

	id, err := repo.CreateMessage("session1", "user123", "Hello")
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	if id != 1 {
		t.Fatalf("Expected ID 1, got %d", id)
	}
}

func TestMessageRepository_GetMessageByID(t *testing.T) {
	repo := createTestMessageRepo(t)

	repo.CreateMessage("session1", "user123", "Hello")

	message, err := repo.GetMessageByID(1)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	if message == nil {
		t.Fatal("Expected message to be found")
	}

	if message.Content != "Hello" {
		t.Fatalf("Expected content 'Hello', got %s", message.Content)
	}
}

func TestMessageRepository_GetMessageByIDNotFound(t *testing.T) {
	repo := createTestMessageRepo(t)

	message, err := repo.GetMessageByID(999)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	if message != nil {
		t.Fatal("Expected nil for non-existent message")
	}
}

func TestMessageRepository_GetMessagesBySession(t *testing.T) {
	repo := createTestMessageRepo(t)

	repo.CreateMessage("session1", "user123", "Message 1")
	repo.CreateMessage("session1", "user456", "Message 2")

	messages, err := repo.GetMessagesBySession("session1", 10, 0)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	if len(messages) != 2 {
		t.Fatalf("Expected 2 messages, got %d", len(messages))
	}
}

func TestMessageRepository_GetMessagesByCursor(t *testing.T) {
	repo := createTestMessageRepo(t)

	repo.CreateMessage("session1", "user123", "Message 1")
	repo.CreateMessage("session1", "user456", "Message 2")
	repo.CreateMessage("session1", "user789", "Message 3")

	messages, err := repo.GetMessagesByCursor("session1", 2, 10)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	if len(messages) != 1 {
		t.Fatalf("Expected 1 message, got %d", len(messages))
	}
}

func TestMessageRepository_UpdateMessage(t *testing.T) {
	repo := createTestMessageRepo(t)

	repo.CreateMessage("session1", "user123", "Original")

	err := repo.UpdateMessage(1, "Updated")
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	message, _ := repo.GetMessageByID(1)
	if message.Content != "Updated" {
		t.Fatalf("Expected content 'Updated', got %s", message.Content)
	}
}

func TestMessageRepository_DeleteMessage(t *testing.T) {
	repo := createTestMessageRepo(t)

	repo.CreateMessage("session1", "user123", "Message")

	err := repo.DeleteMessage(1)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	message, _ := repo.GetMessageByID(1)
	if message != nil {
		t.Fatal("Expected message to be deleted")
	}
}
