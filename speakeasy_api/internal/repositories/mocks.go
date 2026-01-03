package repositories

type CustomError struct {
	Message string
}

func (e *CustomError) Error() string {
	return e.Message
}

type MockUserRepository struct {
	GetUserFunc           func(id string) (*User, error)
	CreateUserFunc        func(id, username string, email, avatarURL *string) error
	UpdateUserFunc        func(id, username string, email, avatarURL *string) error
	DeleteUserFunc        func(id string) error
	GetUserByUsernameFunc func(username string) (*User, error)
}

func (m *MockUserRepository) GetUser(id string) (*User, error) {
	if m.GetUserFunc != nil {
		return m.GetUserFunc(id)
	}
	return nil, nil
}

func (m *MockUserRepository) GetUserByUsername(username string) (*User, error) {
	if m.GetUserByUsernameFunc != nil {
		return m.GetUserByUsernameFunc(username)
	}
	return nil, nil
}

func (m *MockUserRepository) CreateUser(id, username string, email, avatarURL *string) error {
	if m.CreateUserFunc != nil {
		return m.CreateUserFunc(id, username, email, avatarURL)
	}
	return nil
}

func (m *MockUserRepository) UpdateUser(id, username string, email, avatarURL *string) error {
	if m.UpdateUserFunc != nil {
		return m.UpdateUserFunc(id, username, email, avatarURL)
	}
	return nil
}

func (m *MockUserRepository) DeleteUser(id string) error {
	if m.DeleteUserFunc != nil {
		return m.DeleteUserFunc(id)
	}
	return nil
}

type MockSessionRepository struct {
	GetSessionFunc          func(location string) (*Session, error)
	CreateSessionFunc       func(location, status string) error
	UpdateSessionStatusFunc func(location, status string) error
	DeleteSessionFunc       func(location string) error
}

func (m *MockSessionRepository) GetSession(location string) (*Session, error) {
	if m.GetSessionFunc != nil {
		return m.GetSessionFunc(location)
	}
	return nil, nil
}

func (m *MockSessionRepository) CreateSession(location, status string) error {
	if m.CreateSessionFunc != nil {
		return m.CreateSessionFunc(location, status)
	}
	return nil
}

func (m *MockSessionRepository) UpdateSessionStatus(location, status string) error {
	if m.UpdateSessionStatusFunc != nil {
		return m.UpdateSessionStatusFunc(location, status)
	}
	return nil
}

func (m *MockSessionRepository) DeleteSession(location string) error {
	if m.DeleteSessionFunc != nil {
		return m.DeleteSessionFunc(location)
	}
	return nil
}

type MockMessageRepository struct {
	CreateMessageFunc        func(sessionLocation, userID, content string) (int64, error)
	GetMessagesBySessionFunc func(sessionLocation string, limit int, offset int64) ([]Message, error)
	GetMessagesByCursorFunc  func(sessionLocation string, cursorID int64, limit int) ([]Message, error)
	GetMessageByIDFunc       func(id int64) (*Message, error)
	UpdateMessageFunc        func(id int64, content string) error
	DeleteMessageFunc        func(id int64) error
}

func (m *MockMessageRepository) CreateMessage(sessionLocation, userID, content string) (int64, error) {
	if m.CreateMessageFunc != nil {
		return m.CreateMessageFunc(sessionLocation, userID, content)
	}
	return 0, nil
}

func (m *MockMessageRepository) GetMessagesBySession(sessionLocation string, limit int, offset int64) ([]Message, error) {
	if m.GetMessagesBySessionFunc != nil {
		return m.GetMessagesBySessionFunc(sessionLocation, limit, offset)
	}
	return nil, nil
}

func (m *MockMessageRepository) GetMessagesByCursor(sessionLocation string, cursorID int64, limit int) ([]Message, error) {
	if m.GetMessagesByCursorFunc != nil {
		return m.GetMessagesByCursorFunc(sessionLocation, cursorID, limit)
	}
	return nil, nil
}

func (m *MockMessageRepository) GetMessageByID(id int64) (*Message, error) {
	if m.GetMessageByIDFunc != nil {
		return m.GetMessageByIDFunc(id)
	}
	return nil, nil
}

func (m *MockMessageRepository) UpdateMessage(id int64, content string) error {
	if m.UpdateMessageFunc != nil {
		return m.UpdateMessageFunc(id, content)
	}
	return nil
}

func (m *MockMessageRepository) DeleteMessage(id int64) error {
	if m.DeleteMessageFunc != nil {
		return m.DeleteMessageFunc(id)
	}
	return nil
}

type MockSessionUserRepository struct {
	CreateSessionUserFunc       func(sessionID, userID string) error
	UpdateUserLeftTimeFunc      func(sessionID, userID string) error
	GetActiveUsersInSessionFunc func(sessionID string) ([]*SessionUser, error)
	GetActiveUserCountFunc      func(sessionID string) (int, error)
	IsUserInSessionFunc         func(sessionID, userID string) (bool, error)
	GetActiveSessionsFunc       func(userID string) ([]*SessionUser, error)
	JoinSessionWithLockFunc     func(sessionID, userID string) error
	LeaveSessionWithCleanupFunc func(sessionID, userID string) error
}

func (m *MockSessionUserRepository) CreateSessionUser(sessionID, userID string) error {
	if m.CreateSessionUserFunc != nil {
		return m.CreateSessionUserFunc(sessionID, userID)
	}
	return nil
}

func (m *MockSessionUserRepository) UpdateUserLeftTime(sessionID, userID string) error {
	if m.UpdateUserLeftTimeFunc != nil {
		return m.UpdateUserLeftTimeFunc(sessionID, userID)
	}
	return nil
}

func (m *MockSessionUserRepository) GetActiveUsersInSession(sessionID string) ([]*SessionUser, error) {
	if m.GetActiveUsersInSessionFunc != nil {
		return m.GetActiveUsersInSessionFunc(sessionID)
	}
	return nil, nil
}

func (m *MockSessionUserRepository) GetActiveUserCount(sessionID string) (int, error) {
	if m.GetActiveUserCountFunc != nil {
		return m.GetActiveUserCountFunc(sessionID)
	}
	return 0, nil
}

func (m *MockSessionUserRepository) IsUserInSession(sessionID, userID string) (bool, error) {
	if m.IsUserInSessionFunc != nil {
		return m.IsUserInSessionFunc(sessionID, userID)
	}
	return false, nil
}

func (m *MockSessionUserRepository) GetActiveSessions(userID string) ([]*SessionUser, error) {
	if m.GetActiveSessionsFunc != nil {
		return m.GetActiveSessionsFunc(userID)
	}
	return nil, nil
}

func (m *MockSessionUserRepository) JoinSessionWithLock(sessionID, userID string) error {
	if m.JoinSessionWithLockFunc != nil {
		return m.JoinSessionWithLockFunc(sessionID, userID)
	}
	return nil
}

func (m *MockSessionUserRepository) LeaveSessionWithCleanup(sessionID, userID string) error {
	if m.LeaveSessionWithCleanupFunc != nil {
		return m.LeaveSessionWithCleanupFunc(sessionID, userID)
	}
	return nil
}
