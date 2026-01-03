package services

import "github.com/speakeasy/speakeasy-api/internal/repositories"

type UserService struct {
	userRepo repositories.IUserRepository
}

func NewUserService(userRepo repositories.IUserRepository) *UserService {
	return &UserService{
		userRepo: userRepo,
	}
}

func (s *UserService) GetUserByID(userID string) (map[string]interface{}, error) {
	user, err := s.userRepo.GetUser(userID)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, nil
	}
	return map[string]interface{}{
		"id":       user.ID,
		"username": user.Username,
		"email":    user.Email,
		"avatar":   user.AvatarURL,
	}, nil
}

func (s *UserService) CreateUser(username, email, password string) (map[string]interface{}, error) {
	// TODO: Implement CreateUser
	return nil, ErrNotImplemented
}

func (s *UserService) UpdateUser(userID string, data map[string]interface{}) error {
	// TODO: Implement UpdateUser
	return ErrNotImplemented
}

func (s *UserService) DeleteUser(userID string) error {
	return s.userRepo.DeleteUser(userID)
}

func (s *UserService) GetUserPreferences(userID string) (map[string]interface{}, error) {
	// TODO: Implement GetUserPreferences
	return nil, ErrNotImplemented
}

func (s *UserService) UpdateUserPreferences(userID string, preferences map[string]interface{}) error {
	// TODO: Implement UpdateUserPreferences
	return ErrNotImplemented
}
