package repositories

import "context"

type ISessionLocker interface {
	LockSession(ctx context.Context, tx interface{}, sessionID string) error
}

type IUserRepository interface {
	GetUser(id string) (*User, error)
	GetUserByUsername(username string) (*User, error)
	CreateUser(id, username string, email, avatarURL *string) error
	UpdateUser(id, username string, email, avatarURL *string) error
	DeleteUser(id string) error
}

type ISessionRepository interface {
	GetSession(location string) (*Session, error)
	CreateSession(location, status string) error
	UpdateSessionStatus(location, status string) error
	DeleteSession(location string) error
}

type ISessionUserRepository interface {
	CreateSessionUser(sessionID, userID string) error
	UpdateUserLeftTime(sessionID, userID string) error
	GetActiveUsersInSession(sessionID string) ([]*SessionUser, error)
	GetActiveUserCount(sessionID string) (int, error)
	IsUserInSession(sessionID, userID string) (bool, error)
	GetActiveSessions(userID string) ([]*SessionUser, error)
	JoinSessionWithLock(sessionID, userID string) error
	LeaveSessionWithCleanup(sessionID, userID string) error
}

type IMessageRepository interface {
	CreateMessage(sessionLocation, userID, content string) (int64, error)
	GetMessagesBySession(sessionLocation string, limit int, offset int64) ([]Message, error)
	GetMessagesByCursor(sessionLocation string, cursorID int64, limit int) ([]Message, error)
	GetMessageByID(id int64) (*Message, error)
	UpdateMessage(id int64, content string) error
	DeleteMessage(id int64) error
}
