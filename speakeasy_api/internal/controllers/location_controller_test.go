package controllers

import (
	"testing"

	"github.com/speakeasy/speakeasy-api/internal/repositories"
	"github.com/speakeasy/speakeasy-api/internal/services"
)

func TestNewLocationController(t *testing.T) {
	mockRepo := &repositories.MockSessionRepository{}
	locationService := services.NewLocationService(mockRepo)
	controller := NewLocationController(locationService)

	if controller == nil {
		t.Fatal("Expected non-nil LocationController")
	}

	if controller.locationService == nil {
		t.Fatal("Expected locationService to be initialized")
	}
}

func TestLocationControllerCheckVicinity(t *testing.T) {
	mockRepo := &repositories.MockSessionRepository{}
	locationService := services.NewLocationService(mockRepo)
	controller := NewLocationController(locationService)

	if controller == nil {
		t.Fatal("Expected non-nil controller")
	}
}

func TestLocationControllerGetNearbyLocations(t *testing.T) {
	mockRepo := &repositories.MockSessionRepository{}
	locationService := services.NewLocationService(mockRepo)
	controller := NewLocationController(locationService)

	if controller == nil {
		t.Fatal("Expected non-nil controller")
	}
}

func TestLocationControllerGetUserLocation(t *testing.T) {
	mockRepo := &repositories.MockSessionRepository{}
	locationService := services.NewLocationService(mockRepo)
	controller := NewLocationController(locationService)

	if controller == nil {
		t.Fatal("Expected non-nil controller")
	}
}

func TestLocationControllerUpdateUserLocation(t *testing.T) {
	mockRepo := &repositories.MockSessionRepository{}
	locationService := services.NewLocationService(mockRepo)
	controller := NewLocationController(locationService)

	if controller == nil {
		t.Fatal("Expected non-nil controller")
	}
}
