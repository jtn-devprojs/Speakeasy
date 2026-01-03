package repositories

import "database/sql"

type MessageRepository struct {
	db *sql.DB
}

func NewMessageRepository(db *sql.DB) *MessageRepository {
	return &MessageRepository{db: db}
}

type Message struct {
	ID              int64
	SessionLocation string
	UserID          string
	Content         string
	CreatedAt       string
	EditedAt        *string
}

func (r *MessageRepository) CreateMessage(sessionLocation, userID, content string) (int64, error) {
	query := "INSERT INTO messages (session_location, user_id, content) VALUES (?, ?, ?)"
	result, err := r.db.Exec(query, sessionLocation, userID, content)
	if err != nil {
		return 0, err
	}
	return result.LastInsertId()
}

func (r *MessageRepository) GetMessagesBySession(sessionLocation string, limit int, offset int64) ([]Message, error) {
	query := "SELECT id, session_location, user_id, content, created_at, edited_at FROM messages WHERE session_location = ? ORDER BY id DESC LIMIT ? OFFSET ?"
	rows, err := r.db.Query(query, sessionLocation, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var messages []Message
	for rows.Next() {
		var msg Message
		err := rows.Scan(&msg.ID, &msg.SessionLocation, &msg.UserID, &msg.Content, &msg.CreatedAt, &msg.EditedAt)
		if err != nil {
			return nil, err
		}
		messages = append(messages, msg)
	}

	return messages, rows.Err()
}

func (r *MessageRepository) GetMessagesByCursor(sessionLocation string, cursorID int64, limit int) ([]Message, error) {
	query := "SELECT id, session_location, user_id, content, created_at, edited_at FROM messages WHERE session_location = ? AND id < ? ORDER BY id DESC LIMIT ?"
	rows, err := r.db.Query(query, sessionLocation, cursorID, limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var messages []Message
	for rows.Next() {
		var msg Message
		err := rows.Scan(&msg.ID, &msg.SessionLocation, &msg.UserID, &msg.Content, &msg.CreatedAt, &msg.EditedAt)
		if err != nil {
			return nil, err
		}
		messages = append(messages, msg)
	}

	return messages, rows.Err()
}

func (r *MessageRepository) GetMessageByID(id int64) (*Message, error) {
	query := "SELECT id, session_location, user_id, content, created_at, edited_at FROM messages WHERE id = ?"
	row := r.db.QueryRow(query, id)

	var msg Message
	err := row.Scan(&msg.ID, &msg.SessionLocation, &msg.UserID, &msg.Content, &msg.CreatedAt, &msg.EditedAt)
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}
	if err == sql.ErrNoRows {
		return nil, nil
	}

	return &msg, nil
}

func (r *MessageRepository) UpdateMessage(id int64, content string) error {
	query := "UPDATE messages SET content = ?, edited_at = CURRENT_TIMESTAMP WHERE id = ?"
	_, err := r.db.Exec(query, content, id)
	return err
}

func (r *MessageRepository) DeleteMessage(id int64) error {
	query := "DELETE FROM messages WHERE id = ?"
	_, err := r.db.Exec(query, id)
	return err
}
