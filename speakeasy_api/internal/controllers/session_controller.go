package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/speakeasy/speakeasy-api/internal/services"
)

// SessionController handles HTTP requests related to session location
type SessionController struct {
	sessionService services.ISessionService
}

// NewSessionController creates and returns a new SessionController
func NewSessionController(sessionService services.ISessionService) *SessionController {
	return &SessionController{
		sessionService: sessionService,
	}
}

// CheckVicinity checks if a user is within a certain vicinity of a location
func (c *SessionController) CheckVicinity(ctx *gin.Context) {
	// TODO: Implement CheckVicinity
	ctx.JSON(http.StatusNotImplemented, gin.H{"error": "not implemented"})
}

// GetClosestSession retrieves the location of the closest session near the user's current position
func (c *SessionController) GetClosestSession(ctx *gin.Context) {
	// TODO: Implement GetClosestSession
	ctx.JSON(http.StatusNotImplemented, gin.H{"error": "not implemented"})
}

// GetNearbyLocations retrieves locations near the user's current position
func (c *SessionController) GetNearbyLocations(ctx *gin.Context) {
	// TODO: Implement GetNearbyLocations
	ctx.JSON(http.StatusNotImplemented, gin.H{"error": "not implemented"})
}

// GetUserLocation retrieves the current location of the user
func (c *SessionController) GetUserLocation(ctx *gin.Context) {
	// TODO: Implement GetUserLocation
	ctx.JSON(http.StatusNotImplemented, gin.H{"error": "not implemented"})
}

// GetSessionLocations retrieves locations of all sessions
func (c *SessionController) GetSessionLocations(ctx *gin.Context) {
	// TODO: Implement GetSessionLocations
	ctx.JSON(http.StatusNotImplemented, gin.H{"error": "not implemented"})
}

// UpdateUserLocation updates the user's current location
func (c *SessionController) UpdateUserLocation(ctx *gin.Context) {
	// TODO: Implement UpdateUserLocation
	ctx.JSON(http.StatusNotImplemented, gin.H{"error": "not implemented"})
}
