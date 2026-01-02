package services

import "github.com/speakeasy/speakeasy-api/internal/repositories"

type LocationService struct {
	sessionRepo repositories.ISessionRepository
}

func NewLocationService(sessionRepo repositories.ISessionRepository) *LocationService {
	return &LocationService{
		sessionRepo: sessionRepo,
	}
}

func (s *LocationService) IsUserInVicinity(userLat, userLon, targetLat, targetLon, radiusKm float64) bool {
	return false
}

func (s *LocationService) GetNearbyLocations(userLat, userLon, radiusKm float64) (interface{}, error) {
	return nil, nil
}

func (s *LocationService) GetUserLocation(userID string) (interface{}, error) {
	return nil, nil
}

func (s *LocationService) UpdateUserLocation(userID string, lat, lon float64) error {
	return nil
}
