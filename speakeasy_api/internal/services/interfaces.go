package services

type IAuthService interface {
	Logout(token string) error
	ValidateToken(token string) (string, error)
	RefreshToken(token string) (string, error)
}

type ISessionService interface {
	IsUserInVicinity(userLat, userLon, targetLat, targetLon, radiusKm float64) bool
	GetNearbyLocations(userLat, userLon, radiusKm float64) (interface{}, error)
	GetUserLocation(userID string) (interface{}, error)
	UpdateUserLocation(userID string, lat, lon float64) error
}
