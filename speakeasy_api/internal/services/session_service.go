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

func (s *SessionService) JoinSession(sessionID, userID string) error {
	return s.sessionUserRepo.JoinSessionWithLock(sessionID, userID)
}

func (s *SessionService) LeaveSession(sessionID, userID string) error {
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
