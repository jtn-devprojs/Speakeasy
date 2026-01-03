package services

import "github.com/speakeasy/speakeasy-api/internal/repositories"

type AuthService struct {
	userRepo repositories.IUserRepository
}

func NewAuthService(userRepo repositories.IUserRepository) *AuthService {
	return &AuthService{
		userRepo: userRepo,
	}
}

func (s *AuthService) Login(username, password string) (string, error) {
	// TODO: Implement Login
	return "", ErrNotImplemented
}

func (s *AuthService) Logout(token string) error {
	// TODO: Implement Logout
	return ErrNotImplemented
}

func (s *AuthService) Register(username, email, password string) (map[string]interface{}, error) {
	// TODO: Implement Register
	return nil, ErrNotImplemented
}

func (s *AuthService) ValidateToken(token string) (string, error) {
	// TODO: Implement ValidateToken
	return "", ErrNotImplemented
}

func (s *AuthService) RefreshToken(token string) (string, error) {
	// TODO: Implement RefreshToken
	return "", ErrNotImplemented
}
