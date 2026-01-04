package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/speakeasy/speakeasy-api/internal/services"
)

type SessionController struct {
	sessionService services.ISessionService
}

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

func (c *SessionController) GetClosestSession(ctx *gin.Context) {
	// TODO: Implement GetClosestSession
	ctx.JSON(http.StatusNotImplemented, gin.H{"error": "not implemented"})
}

func (c *SessionController) GetNearbyLocations(ctx *gin.Context) {
	// TODO: Implement GetNearbyLocations
	ctx.JSON(http.StatusNotImplemented, gin.H{"error": "not implemented"})
}

func (c *SessionController) GetUserLocation(ctx *gin.Context) {
	// TODO: Implement GetUserLocation
	ctx.JSON(http.StatusNotImplemented, gin.H{"error": "not implemented"})
}

func (c *SessionController) UpdateUserLocation(ctx *gin.Context) {
	// TODO: Implement UpdateUserLocation
	ctx.JSON(http.StatusNotImplemented, gin.H{"error": "not implemented"})
}
