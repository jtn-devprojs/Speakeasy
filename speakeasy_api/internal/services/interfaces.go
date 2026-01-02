package services

type IUserService interface {
	GetUserByID(userID string) (map[string]interface{}, error)
	CreateUser(username, email, password string) (map[string]interface{}, error)
	UpdateUser(userID string, data map[string]interface{}) error
	DeleteUser(userID string) error
	GetUserPreferences(userID string) (map[string]interface{}, error)
	UpdateUserPreferences(userID string, preferences map[string]interface{}) error
}

type IAuthService interface {
	Login(username, password string) (string, error)
	Logout(token string) error
	Register(username, email, password string) (map[string]interface{}, error)
	ValidateToken(token string) (string, error)
	RefreshToken(token string) (string, error)
}

type ILocationService interface {
	IsUserInVicinity(userLat, userLon, targetLat, targetLon, radiusKm float64) bool
	GetNearbyLocations(userLat, userLon, radiusKm float64) (interface{}, error)
	GetUserLocation(userID string) (interface{}, error)
	UpdateUserLocation(userID string, lat, lon float64) error
}
