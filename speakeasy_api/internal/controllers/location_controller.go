package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/speakeasy/speakeasy-api/internal/services"
)

// LocationController handles HTTP requests related to user location
type LocationController struct {
	locationService services.ILocationService
}

// NewLocationController creates and returns a new LocationController
func NewLocationController(locationService services.ILocationService) *LocationController {
	return &LocationController{
		locationService: locationService,
	}
}

// CheckVicinity checks if a user is within a certain vicinity of a location
func (c *LocationController) CheckVicinity(ctx *gin.Context) {
	// TODO: Implement CheckVicinity
	ctx.JSON(http.StatusNotImplemented, gin.H{"error": "not implemented"})
}

// GetNearbyLocations retrieves locations near the user's current position
func (c *LocationController) GetNearbyLocations(ctx *gin.Context) {
	// TODO: Implement GetNearbyLocations
	ctx.JSON(http.StatusNotImplemented, gin.H{"error": "not implemented"})
}

// GetUserLocation retrieves the current location of the user
func (c *LocationController) GetUserLocation(ctx *gin.Context) {
	// TODO: Implement GetUserLocation
	ctx.JSON(http.StatusNotImplemented, gin.H{"error": "not implemented"})
}

// UpdateUserLocation updates the user's current location
func (c *LocationController) UpdateUserLocation(ctx *gin.Context) {
	// TODO: Implement UpdateUserLocation
	ctx.JSON(http.StatusNotImplemented, gin.H{"error": "not implemented"})
}
