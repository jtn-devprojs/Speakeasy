package services

type MockAuthService struct {
	LogoutFunc        func(token string) error
	ValidateTokenFunc func(token string) (string, error)
	RefreshTokenFunc  func(token string) (string, error)
}

func (m *MockAuthService) Logout(token string) error {
	if m.LogoutFunc != nil {
		return m.LogoutFunc(token)
	}
	return nil
}

func (m *MockAuthService) ValidateToken(token string) (string, error) {
	if m.ValidateTokenFunc != nil {
		return m.ValidateTokenFunc(token)
	}
	return "", ErrNotImplemented
}

func (m *MockAuthService) RefreshToken(token string) (string, error) {
	if m.RefreshTokenFunc != nil {
		return m.RefreshTokenFunc(token)
	}
	return "", ErrNotImplemented
}

type MockSessionService struct {
	IsUserInVicinityFunc   func(userLat, userLon, targetLat, targetLon, radiusKm float64) bool
	GetNearbyLocationsFunc func(userLat, userLon, radiusKm float64) (interface{}, error)
	GetUserLocationFunc    func(userID string) (interface{}, error)
	UpdateUserLocationFunc func(userID string, lat, lon float64) error
}

func (m *MockSessionService) IsUserInVicinity(userLat, userLon, targetLat, targetLon, radiusKm float64) bool {
	if m.IsUserInVicinityFunc != nil {
		return m.IsUserInVicinityFunc(userLat, userLon, targetLat, targetLon, radiusKm)
	}
	return false
}

func (m *MockSessionService) GetNearbyLocations(userLat, userLon, radiusKm float64) (interface{}, error) {
	if m.GetNearbyLocationsFunc != nil {
		return m.GetNearbyLocationsFunc(userLat, userLon, radiusKm)
	}
	return nil, ErrNotImplemented
}

func (m *MockSessionService) GetUserLocation(userID string) (interface{}, error) {
	if m.GetUserLocationFunc != nil {
		return m.GetUserLocationFunc(userID)
	}
	return nil, ErrNotImplemented
}

func (m *MockSessionService) UpdateUserLocation(userID string, lat, lon float64) error {
	if m.UpdateUserLocationFunc != nil {
		return m.UpdateUserLocationFunc(userID, lat, lon)
	}
	return ErrNotImplemented
}
