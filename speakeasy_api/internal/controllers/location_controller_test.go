package controllers

import (
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/speakeasy/speakeasy-api/internal/services"
)

func TestNewLocationController(t *testing.T) {
	locationService := services.NewLocationService()
	controller := NewLocationController(locationService)

	if controller == nil {
		t.Fatal("Expected non-nil LocationController")
	}

	if controller.locationService == nil {
		t.Fatal("Expected locationService to be initialized")
	}
}

func TestLocationControllerCheckVicinity(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.New()
	locationService := services.NewLocationService()
	controller := NewLocationController(locationService)

	router.POST("/check-vicinity", controller.CheckVicinity)

	// TODO: Implement test logic when CheckVicinity is implemented
	// For now, this is a stub to verify the controller is wired correctly
}

func TestLocationControllerGetNearbyLocations(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.New()
	locationService := services.NewLocationService()
	controller := NewLocationController(locationService)

	router.GET("/nearby", controller.GetNearbyLocations)

	// TODO: Implement test logic when GetNearbyLocations is implemented
	// For now, this is a stub to verify the controller is wired correctly
}

func TestLocationControllerGetUserLocation(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.New()
	locationService := services.NewLocationService()
	controller := NewLocationController(locationService)

	router.GET("/location", controller.GetUserLocation)

	// TODO: Implement test logic when GetUserLocation is implemented
	// For now, this is a stub to verify the controller is wired correctly
}

func TestLocationControllerUpdateUserLocation(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.New()
	locationService := services.NewLocationService()
	controller := NewLocationController(locationService)

	router.PUT("/location", controller.UpdateUserLocation)

	// TODO: Implement test logic when UpdateUserLocation is implemented
	// For now, this is a stub to verify the controller is wired correctly
}
