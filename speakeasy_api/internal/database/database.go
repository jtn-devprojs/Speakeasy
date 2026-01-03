package database

import (
	"database/sql"
	"fmt"

	_ "modernc.org/sqlite"
)

func Init(dbType, connection string) (*sql.DB, error) {
	switch dbType {
	case "sqlite":
		return initSQLite(connection)
	case "mysql":
		return initMySQL(connection)
	default:
		return nil, fmt.Errorf("unsupported database type: %s", dbType)
	}
}

func initSQLite(connection string) (*sql.DB, error) {
	db, err := sql.Open("sqlite", connection)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	if err := createTables(db); err != nil {
		return nil, err
	}

	return db, nil
}

func initMySQL(connection string) (*sql.DB, error) {
	db, err := sql.Open("mysql", connection)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	if err := createTables(db); err != nil {
		return nil, err
	}

	return db, nil
}

func createTables(db *sql.DB) error {
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

	_, err := db.Exec(schema)
	return err
}
