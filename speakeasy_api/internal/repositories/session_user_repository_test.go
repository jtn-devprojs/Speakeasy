package repositories

import (
	"database/sql"
	"testing"

	_ "modernc.org/sqlite"
)

func TestNewSessionUserRepository(t *testing.T) {
	db := createSessionUserTestDB(t)
	defer db.Close()

	repo := NewSessionUserRepository(db)

	if repo == nil {
		t.Fatal("Expected non-nil SessionUserRepository")
	}
}

func TestSessionUserRepository_CreateSessionUser(t *testing.T) {
	db := createSessionUserTestDB(t)
	defer db.Close()

	repo := NewSessionUserRepository(db)

	err := repo.CreateSessionUser("session-1", "user-1")
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	users, err := repo.GetActiveUsersInSession("session-1")
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if len(users) != 1 {
		t.Fatalf("Expected 1 user, got %d", len(users))
	}

	if users[0].UserID != "user-1" {
		t.Fatalf("Expected user-1, got %s", users[0].UserID)
	}

	if users[0].LeftAt != nil {
		t.Fatal("Expected LeftAt to be nil for active user")
	}
}

func TestSessionUserRepository_UpdateUserLeftTime(t *testing.T) {
	db := createSessionUserTestDB(t)
	defer db.Close()

	repo := NewSessionUserRepository(db)

	err := repo.CreateSessionUser("session-1", "user-1")
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	err = repo.UpdateUserLeftTime("session-1", "user-1")
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	users, err := repo.GetActiveUsersInSession("session-1")
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if len(users) != 0 {
		t.Fatalf("Expected 0 users, got %d", len(users))
	}
}

func TestSessionUserRepository_GetActiveUsersInSession(t *testing.T) {
	db := createSessionUserTestDB(t)
	defer db.Close()

	repo := NewSessionUserRepository(db)

	repo.CreateSessionUser("session-1", "user-1")
	repo.CreateSessionUser("session-1", "user-2")
	repo.CreateSessionUser("session-1", "user-3")

	users, err := repo.GetActiveUsersInSession("session-1")
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if len(users) != 3 {
		t.Fatalf("Expected 3 users, got %d", len(users))
	}
}

func TestSessionUserRepository_GetActiveUserCount(t *testing.T) {
	db := createSessionUserTestDB(t)
	defer db.Close()

	repo := NewSessionUserRepository(db)

	repo.CreateSessionUser("session-1", "user-1")
	repo.CreateSessionUser("session-1", "user-2")
	repo.CreateSessionUser("session-1", "user-3")

	count, err := repo.GetActiveUserCount("session-1")
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if count != 3 {
		t.Fatalf("Expected count 3, got %d", count)
	}

	repo.UpdateUserLeftTime("session-1", "user-1")

	count, err = repo.GetActiveUserCount("session-1")
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if count != 2 {
		t.Fatalf("Expected count 2 after user left, got %d", count)
	}
}

func TestSessionUserRepository_IsUserInSession(t *testing.T) {
	db := createSessionUserTestDB(t)
	defer db.Close()

	repo := NewSessionUserRepository(db)

	repo.CreateSessionUser("session-1", "user-1")

	inSession, err := repo.IsUserInSession("session-1", "user-1")
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if !inSession {
		t.Fatal("Expected user to be in session")
	}

	inSession, err = repo.IsUserInSession("session-1", "user-999")
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if inSession {
		t.Fatal("Expected user-999 to not be in session")
	}

	repo.UpdateUserLeftTime("session-1", "user-1")

	inSession, err = repo.IsUserInSession("session-1", "user-1")
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if inSession {
		t.Fatal("Expected user-1 to not be in session after leaving")
	}
}

func TestSessionUserRepository_GetActiveSessions(t *testing.T) {
	db := createSessionUserTestDB(t)
	defer db.Close()

	repo := NewSessionUserRepository(db)

	repo.CreateSessionUser("session-1", "user-1")
	repo.CreateSessionUser("session-2", "user-1")
	repo.CreateSessionUser("session-3", "user-1")
	repo.CreateSessionUser("session-4", "user-2")

	sessions, err := repo.GetActiveSessions("user-1")
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if len(sessions) != 3 {
		t.Fatalf("Expected 3 sessions, got %d", len(sessions))
	}

	sessions, err = repo.GetActiveSessions("user-2")
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if len(sessions) != 1 {
		t.Fatalf("Expected 1 session, got %d", len(sessions))
	}
}

func TestSessionUserRepository_DuplicateUserInSession(t *testing.T) {
	db := createSessionUserTestDB(t)
	defer db.Close()

	repo := NewSessionUserRepository(db)

	err := repo.CreateSessionUser("session-1", "user-1")
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	err = repo.CreateSessionUser("session-1", "user-1")
	if err == nil {
		t.Fatal("Expected error when adding duplicate user to session")
	}
}

func TestSessionUserRepository_JoinSessionWithLock(t *testing.T) {
	db := createSessionUserTestDB(t)
	defer db.Close()

	// Setup: Create a session in sessions table
	_, err := db.Exec(
		"INSERT INTO sessions (id, latitude, longitude, accuracy) VALUES (?, ?, ?, ?)",
		"session-1", 40.7128, -74.0060, 100,
	)
	if err != nil {
		t.Fatalf("Failed to create session: %v", err)
	}

	repo := NewSessionUserRepository(db)

	// Test successful join
	err = repo.JoinSessionWithLock("session-1", "user-1")
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	// Verify user is now in session
	inSession, err := repo.IsUserInSession("session-1", "user-1")
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if !inSession {
		t.Fatal("Expected user to be in session after join")
	}
}

func TestSessionUserRepository_JoinSessionWithLock_SessionNotFound(t *testing.T) {
	db := createSessionUserTestDB(t)
	defer db.Close()

	repo := NewSessionUserRepository(db)

	// Try to join non-existent session
	err := repo.JoinSessionWithLock("non-existent", "user-1")
	if err == nil {
		t.Fatal("Expected error when joining non-existent session")
	}

	if err.Error() != "session not found" {
		t.Fatalf("Expected 'session not found' error, got %v", err)
	}
}

func TestSessionUserRepository_JoinSessionWithLock_EndedSession(t *testing.T) {
	db := createSessionUserTestDB(t)
	defer db.Close()

	// Setup: Create an ended session
	_, err := db.Exec(
		"INSERT INTO sessions (id, latitude, longitude, accuracy, ended_at) VALUES (?, ?, ?, ?, CURRENT_TIMESTAMP)",
		"session-ended", 40.7128, -74.0060, 100,
	)
	if err != nil {
		t.Fatalf("Failed to create ended session: %v", err)
	}

	repo := NewSessionUserRepository(db)

	// Try to join ended session
	err = repo.JoinSessionWithLock("session-ended", "user-1")
	if err == nil {
		t.Fatal("Expected error when joining ended session")
	}

	if err.Error() != "cannot join ended session" {
		t.Fatalf("Expected 'cannot join ended session' error, got %v", err)
	}
}

func TestSessionUserRepository_JoinSessionWithLock_DuplicateJoin(t *testing.T) {
	db := createSessionUserTestDB(t)
	defer db.Close()

	// Setup: Create a session and add a user
	_, err := db.Exec(
		"INSERT INTO sessions (id, latitude, longitude, accuracy) VALUES (?, ?, ?, ?)",
		"session-1", 40.7128, -74.0060, 100,
	)
	if err != nil {
		t.Fatalf("Failed to create session: %v", err)
	}

	repo := NewSessionUserRepository(db)

	// First join succeeds
	err = repo.JoinSessionWithLock("session-1", "user-1")
	if err != nil {
		t.Fatalf("First join failed: %v", err)
	}

	// Second join should fail
	err = repo.JoinSessionWithLock("session-1", "user-1")
	if err == nil {
		t.Fatal("Expected error when user joins session twice")
	}

	if err.Error() != "user already in session" {
		t.Fatalf("Expected 'user already in session' error, got %v", err)
	}
}

func TestSessionUserRepository_LeaveSessionWithCleanup(t *testing.T) {
	db := createSessionUserTestDB(t)
	defer db.Close()

	// Setup: Create a session with users
	_, err := db.Exec(
		"INSERT INTO sessions (id, latitude, longitude, accuracy) VALUES (?, ?, ?, ?)",
		"session-1", 40.7128, -74.0060, 100,
	)
	if err != nil {
		t.Fatalf("Failed to create session: %v", err)
	}

	repo := NewSessionUserRepository(db)
	repo.JoinSessionWithLock("session-1", "user-1")
	repo.JoinSessionWithLock("session-1", "user-2")

	// User-1 leaves (session should still be active)
	err = repo.LeaveSessionWithCleanup("session-1", "user-1")
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	// Verify user-1 is no longer in session
	inSession, err := repo.IsUserInSession("session-1", "user-1")
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if inSession {
		t.Fatal("Expected user-1 to not be in session after leaving")
	}

	// Verify user-2 is still in session
	inSession, err = repo.IsUserInSession("session-1", "user-2")
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if !inSession {
		t.Fatal("Expected user-2 to still be in session")
	}

	// Verify session is still active (not ended)
	var endedAt sql.NullTime
	err = db.QueryRow("SELECT ended_at FROM sessions WHERE id = ?", "session-1").Scan(&endedAt)
	if err != nil {
		t.Fatalf("Failed to query session: %v", err)
	}

	if endedAt.Valid {
		t.Fatal("Expected session to still be active")
	}
}

func TestSessionUserRepository_LeaveSessionWithCleanup_LastUserLeaves(t *testing.T) {
	db := createSessionUserTestDB(t)
	defer db.Close()

	// Setup: Create a session with one user
	_, err := db.Exec(
		"INSERT INTO sessions (id, latitude, longitude, accuracy) VALUES (?, ?, ?, ?)",
		"session-1", 40.7128, -74.0060, 100,
	)
	if err != nil {
		t.Fatalf("Failed to create session: %v", err)
	}

	repo := NewSessionUserRepository(db)
	repo.JoinSessionWithLock("session-1", "user-1")

	// User-1 leaves (session should be marked as ended)
	err = repo.LeaveSessionWithCleanup("session-1", "user-1")
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	// Verify session is now ended
	var endedAt sql.NullTime
	err = db.QueryRow("SELECT ended_at FROM sessions WHERE id = ?", "session-1").Scan(&endedAt)
	if err != nil {
		t.Fatalf("Failed to query session: %v", err)
	}

	if !endedAt.Valid {
		t.Fatal("Expected session to be marked as ended")
	}
}

func createSessionUserTestDB(t *testing.T) *sql.DB {
	db, err := sql.Open("sqlite", ":memory:")
	if err != nil {
		t.Fatalf("Failed to create test database: %v", err)
	}

	schema := `
	CREATE TABLE sessions (
		id TEXT PRIMARY KEY,
		latitude FLOAT NOT NULL,
		longitude FLOAT NOT NULL,
		accuracy INT NOT NULL,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		ended_at TIMESTAMP NULL
	);

	CREATE TABLE users (
		id TEXT PRIMARY KEY,
		username TEXT NOT NULL UNIQUE,
		email TEXT,
		avatar_url TEXT,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	);

	CREATE TABLE session_users (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		session_id TEXT NOT NULL,
		user_id TEXT NOT NULL,
		joined_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		left_at TIMESTAMP NULL,
		FOREIGN KEY (session_id) REFERENCES sessions(id),
		FOREIGN KEY (user_id) REFERENCES users(id),
		UNIQUE(session_id, user_id)
	);
	`

	_, err = db.Exec(schema)
	if err != nil {
		t.Fatalf("Failed to create schema: %v", err)
	}

	return db
}
