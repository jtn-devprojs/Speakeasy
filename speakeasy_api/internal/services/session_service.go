package services

import "github.com/speakeasy/speakeasy-api/internal/repositories"

type SessionService struct {
	sessionRepo     repositories.ISessionRepository
	sessionUserRepo repositories.ISessionUserRepository
}

func NewSessionService(
	sessionRepo repositories.ISessionRepository,
	sessionUserRepo repositories.ISessionUserRepository,
) *SessionService {
	return &SessionService{
		sessionRepo:     sessionRepo,
		sessionUserRepo: sessionUserRepo,
	}
}

// JoinSession handles user joining an active session
func (s *SessionService) JoinSession(sessionID, userID string) error {
	// Repository handles transaction, lock, and validation
	return s.sessionUserRepo.JoinSessionWithLock(sessionID, userID)
}

// LeaveSession handles user leaving a session
func (s *SessionService) LeaveSession(sessionID, userID string) error {
	// Repository handles transaction and cleanup
	return s.sessionUserRepo.LeaveSessionWithCleanup(sessionID, userID)
}

func (s *SessionService) IsUserInVicinity(userLat, userLon, targetLat, targetLon, radiusKm float64) bool {
	return false
}

func (s *SessionService) GetNearbyLocations(userLat, userLon, radiusKm float64) (interface{}, error) {
	return nil, nil
}

func (s *SessionService) GetUserLocation(userID string) (interface{}, error) {
	return nil, nil
}

func (s *SessionService) UpdateUserLocation(userID string, lat, lon float64) error {
	return nil
}
