package repositories

import "database/sql"

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db: db}
}

type User struct {
	ID        string
	Username  string
	Email     *string
	AvatarURL *string
	CreatedAt string
}

func (r *UserRepository) GetUser(id string) (*User, error) {
	query := "SELECT id, username, email, avatar_url, created_at FROM users WHERE id = ?"
	row := r.db.QueryRow(query, id)

	var user User
	err := row.Scan(&user.ID, &user.Username, &user.Email, &user.AvatarURL, &user.CreatedAt)
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}
	if err == sql.ErrNoRows {
		return nil, nil
	}

	return &user, nil
}

func (r *UserRepository) GetUserByUsername(username string) (*User, error) {
	query := "SELECT id, username, email, avatar_url, created_at FROM users WHERE username = ?"
	row := r.db.QueryRow(query, username)

	var user User
	err := row.Scan(&user.ID, &user.Username, &user.Email, &user.AvatarURL, &user.CreatedAt)
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}
	if err == sql.ErrNoRows {
		return nil, nil
	}

	return &user, nil
}

func (r *UserRepository) CreateUser(id, username string, email, avatarURL *string) error {
	query := "INSERT INTO users (id, username, email, avatar_url) VALUES (?, ?, ?, ?)"
	_, err := r.db.Exec(query, id, username, email, avatarURL)
	return err
}

func (r *UserRepository) UpdateUser(id, username string, email, avatarURL *string) error {
	query := "UPDATE users SET username = ?, email = ?, avatar_url = ? WHERE id = ?"
	_, err := r.db.Exec(query, username, email, avatarURL, id)
	return err
}

func (r *UserRepository) DeleteUser(id string) error {
	query := "DELETE FROM users WHERE id = ?"
	_, err := r.db.Exec(query, id)
	return err
}
