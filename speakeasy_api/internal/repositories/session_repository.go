package repositories

import "database/sql"

type SessionRepository struct {
	db *sql.DB
}

func NewSessionRepository(db *sql.DB) *SessionRepository {
	return &SessionRepository{db: db}
}

type Session struct {
	Location  string
	CreatedAt string
	Status    string
}

func (r *SessionRepository) GetSession(location string) (*Session, error) {
	query := "SELECT location, created_at, status FROM sessions WHERE location = ?"
	row := r.db.QueryRow(query, location)

	var session Session
	err := row.Scan(&session.Location, &session.CreatedAt, &session.Status)
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}
	if err == sql.ErrNoRows {
		return nil, nil
	}

	return &session, nil
}

func (r *SessionRepository) CreateSession(location, status string) error {
	query := "INSERT INTO sessions (location, status) VALUES (?, ?)"
	_, err := r.db.Exec(query, location, status)
	return err
}

func (r *SessionRepository) UpdateSessionStatus(location, status string) error {
	query := "UPDATE sessions SET status = ? WHERE location = ?"
	_, err := r.db.Exec(query, status, location)
	return err
}

func (r *SessionRepository) DeleteSession(location string) error {
	query := "DELETE FROM sessions WHERE location = ?"
	_, err := r.db.Exec(query, location)
	return err
}
