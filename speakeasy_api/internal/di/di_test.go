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

	container := NewContainer(db)

	if container == nil {
		t.Fatal("Expected non-nil container")
	}

	if container.DB == nil {
		t.Fatal("Expected DB to be initialized")
	}

	if container.UserService == nil {
		t.Fatal("Expected UserService to be initialized")
	}

	if container.AuthService == nil {
		t.Fatal("Expected AuthService to be initialized")
	}

	if container.LocationService == nil {
		t.Fatal("Expected LocationService to be initialized")
	}

	if container.UserController == nil {
		t.Fatal("Expected UserController to be initialized")
	}

	if container.AuthController == nil {
		t.Fatal("Expected AuthController to be initialized")
	}

	if container.LocationController == nil {
		t.Fatal("Expected LocationController to be initialized")
	}
}

func TestContainer_DependencyInjection(t *testing.T) {
	db := createTestDB(t)
	defer db.Close()

	container := NewContainer(db)

	if container.UserController == nil {
		t.Fatal("UserController should be initialized")
	}

	if container.AuthController == nil {
		t.Fatal("AuthController should be initialized")
	}

	if container.LocationController == nil {
		t.Fatal("LocationController should be initialized")
	}

	if container.UserService == nil {
		t.Fatal("UserService should be initialized")
	}

	if container.AuthService == nil {
		t.Fatal("AuthService should be initialized")
	}

	if container.LocationService == nil {
		t.Fatal("LocationService should be initialized")
	}
}

func TestContainer_Singleton(t *testing.T) {
	db := createTestDB(t)
	defer db.Close()

	container1 := NewContainer(db)
	container2 := NewContainer(db)

	if container1 == container2 {
		t.Fatal("Expected separate container instances")
	}

	if container1.UserService != container1.UserService {
		t.Fatal("UserService should be consistent within container")
	}
}
