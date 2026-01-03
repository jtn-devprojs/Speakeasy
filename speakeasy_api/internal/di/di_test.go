package di

import (
	"database/sql"
	"testing"

	_ "modernc.org/sqlite"
)

func createTestDB(t *testing.T) *sql.DB {
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

	CREATE TABLE IF NOT EXISTS session_users (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		session_location TEXT NOT NULL,
		user_id TEXT NOT NULL,
		joined_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		is_active BOOLEAN DEFAULT 1,
		FOREIGN KEY (session_location) REFERENCES sessions(location)
	);

	CREATE TABLE IF NOT EXISTS messages (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		session_location TEXT NOT NULL,
		user_id TEXT NOT NULL,
		content TEXT NOT NULL,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		edited_at TIMESTAMP,
		FOREIGN KEY (session_location) REFERENCES sessions(location)
	);

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

	return db
}

func TestNewContainer(t *testing.T) {
	db := createTestDB(t)
	defer db.Close()

	container := NewContainer(db, "sqlite")

	if container == nil {
		t.Fatal("Expected non-nil container")
	}

	if container.DB == nil {
		t.Fatal("Expected DB to be initialized")
	}

	if container.AuthService == nil {
		t.Fatal("Expected AuthService to be initialized")
	}

	if container.SessionService == nil {
		t.Fatal("Expected SessionService to be initialized")
	}

	if container.UserRepo == nil {
		t.Fatal("Expected UserRepo to be initialized")
	}

	if container.SessionRepo == nil {
		t.Fatal("Expected SessionRepo to be initialized")
	}

	if container.SessionUserRepo == nil {
		t.Fatal("Expected SessionUserRepo to be initialized")
	}

	if container.SessionController == nil {
		t.Fatal("Expected SessionController to be initialized")
	}
}

func TestContainer_DependencyInjection(t *testing.T) {
	db := createTestDB(t)
	defer db.Close()

	container := NewContainer(db, "sqlite")

	if container.SessionController == nil {
		t.Fatal("SessionController should be initialized")
	}

	if container.AuthService == nil {
		t.Fatal("AuthService should be initialized")
	}

	if container.SessionService == nil {
		t.Fatal("SessionService should be initialized")
	}
}

func TestContainer_Singleton(t *testing.T) {
	db := createTestDB(t)
	defer db.Close()

	container1 := NewContainer(db, "sqlite")
	container2 := NewContainer(db, "sqlite")

	if container1 == container2 {
		t.Fatal("Expected separate container instances")
	}

	if container1.AuthService != container1.AuthService {
		t.Fatal("AuthService should be consistent within container")
	}
}
