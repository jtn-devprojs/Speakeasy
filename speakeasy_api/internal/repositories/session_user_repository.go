package repositories

import (
	"database/sql"
	"fmt"
	"time"
)

type SessionUserRepository struct {
	db *sql.DB
}

func NewSessionUserRepository(db *sql.DB) ISessionUserRepository {
	return &SessionUserRepository{db: db}
}

type SessionUser struct {
	ID        int
	SessionID string
	UserID    string
	JoinedAt  time.Time
	LeftAt    *time.Time
}

func (r *SessionUserRepository) CreateSessionUser(sessionID, userID string) error {
	query := "INSERT INTO session_users (session_id, user_id, joined_at) VALUES (?, ?, CURRENT_TIMESTAMP)"
	_, err := r.db.Exec(query, sessionID, userID)
	return err
}

func (r *SessionUserRepository) UpdateUserLeftTime(sessionID, userID string) error {
	query := "UPDATE session_users SET left_at = CURRENT_TIMESTAMP WHERE session_id = ? AND user_id = ?"
	_, err := r.db.Exec(query, sessionID, userID)
	return err
}

func (r *SessionUserRepository) GetActiveUsersInSession(sessionID string) ([]*SessionUser, error) {
	query := "SELECT id, session_id, user_id, joined_at, left_at FROM session_users WHERE session_id = ? AND left_at IS NULL"
	rows, err := r.db.Query(query, sessionID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []*SessionUser
	for rows.Next() {
		var user SessionUser
		err := rows.Scan(&user.ID, &user.SessionID, &user.UserID, &user.JoinedAt, &user.LeftAt)
		if err != nil {
			return nil, err
		}
		users = append(users, &user)
	}

	return users, rows.Err()
}

func (r *SessionUserRepository) GetActiveUserCount(sessionID string) (int, error) {
	query := "SELECT COUNT(*) FROM session_users WHERE session_id = ? AND left_at IS NULL"
	var count int
	err := r.db.QueryRow(query, sessionID).Scan(&count)
	if err != nil {
		return 0, err
	}
	return count, nil
}

func (r *SessionUserRepository) IsUserInSession(sessionID, userID string) (bool, error) {
	query := "SELECT COUNT(*) FROM session_users WHERE session_id = ? AND user_id = ? AND left_at IS NULL"
	var count int
	err := r.db.QueryRow(query, sessionID, userID).Scan(&count)
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

func (r *SessionUserRepository) GetActiveSessions(userID string) ([]*SessionUser, error) {
	query := "SELECT id, session_id, user_id, joined_at, left_at FROM session_users WHERE user_id = ? AND left_at IS NULL"
	rows, err := r.db.Query(query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var sessions []*SessionUser
	for rows.Next() {
		var session SessionUser
		err := rows.Scan(&session.ID, &session.SessionID, &session.UserID, &session.JoinedAt, &session.LeftAt)
		if err != nil {
			return nil, err
		}
		sessions = append(sessions, &session)
	}

	return sessions, rows.Err()
}

// JoinSessionWithLock handles user joining a session with transaction and lock
func (r *SessionUserRepository) JoinSessionWithLock(sessionID, userID string) error {
	// Start transaction
	tx, err := r.db.Begin()
	if err != nil {
		return fmt.Errorf("failed to begin transaction: %w", err)
	}
	defer tx.Rollback()

	// Lock the session row to prevent concurrent modifications
	var endedAt sql.NullTime
	err = tx.QueryRow(
		"SELECT ended_at FROM sessions WHERE id = ? FOR UPDATE",
		sessionID,
	).Scan(&endedAt)

	if err == sql.ErrNoRows {
		return fmt.Errorf("session not found")
	}
	if err != nil {
		return fmt.Errorf("failed to lock session: %w", err)
	}

	// Check if session is still active
	if endedAt.Valid {
		return fmt.Errorf("cannot join ended session")
	}

	// Insert user into session
	_, err = tx.Exec(
		"INSERT INTO session_users (session_id, user_id, joined_at) VALUES (?, ?, CURRENT_TIMESTAMP)",
		sessionID, userID,
	)
	if err != nil {
		// Check for duplicate join attempt
		if err.Error() == "UNIQUE constraint failed: session_users.session_id, session_users.user_id" {
			return fmt.Errorf("user already in session")
		}
		return fmt.Errorf("failed to join session: %w", err)
	}

	// Commit transaction (lock released)
	return tx.Commit().Err
}

// LeaveSessionWithCleanup handles user leaving a session and marks session as ended if empty
func (r *SessionUserRepository) LeaveSessionWithCleanup(sessionID, userID string) error {
	// Start transaction
	tx, err := r.db.Begin()
	if err != nil {
		return fmt.Errorf("failed to begin transaction: %w", err)
	}
	defer tx.Rollback()

	// Update user's left_at timestamp
	_, err = tx.Exec(
		"UPDATE session_users SET left_at = CURRENT_TIMESTAMP WHERE session_id = ? AND user_id = ?",
		sessionID, userID,
	)
	if err != nil {
		return fmt.Errorf("failed to update user left time: %w", err)
	}

	// Check if session now has no active users
	var count int
	err = tx.QueryRow(
		"SELECT COUNT(*) FROM session_users WHERE session_id = ? AND left_at IS NULL",
		sessionID,
	).Scan(&count)
	if err != nil {
		return fmt.Errorf("failed to get active user count: %w", err)
	}

	// If session is empty, mark it as ended
	if count == 0 {
		_, err = tx.Exec(
			"UPDATE sessions SET ended_at = CURRENT_TIMESTAMP WHERE id = ?",
			sessionID,
		)
		if err != nil {
			return fmt.Errorf("failed to end session: %w", err)
		}
	}

	// Commit transaction
	return tx.Commit().Err
}
