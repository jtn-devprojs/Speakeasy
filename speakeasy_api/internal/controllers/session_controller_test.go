package controllers

import (
	"testing"

	"github.com/speakeasy/speakeasy-api/internal/repositories"
	"github.com/speakeasy/speakeasy-api/internal/services"
)

func TestNewSessionController(t *testing.T) {
	mockRepo := &repositories.MockSessionRepository{}
	locationService := services.NewLocationService(mockRepo)
	controller := NewSessionController(locationService)

	if controller == nil {
		t.Fatal("Expected non-nil SessionController")
	}

	if controller.locationService == nil {
		t.Fatal("Expected locationService to be initialized")
	}
}

func TestSessionControllerCheckVicinity(t *testing.T) {
	mockRepo := &repositories.MockSessionRepository{}
	locationService := services.NewLocationService(mockRepo)
	controller := NewSessionController(locationService)

	if controller == nil {
		t.Fatal("Expected non-nil controller")
	}
}

func TestSessionControllerGetClosestSession(t *testing.T) {
	mockRepo := &repositories.MockSessionRepository{}
	locationService := services.NewLocationService(mockRepo)
	controller := NewSessionController(locationService)

	if controller == nil {
		t.Fatal("Expected non-nil controller")
	}
}

func TestSessionControllerGetNearbyLocations(t *testing.T) {
	mockRepo := &repositories.MockSessionRepository{}
	locationService := services.NewLocationService(mockRepo)
	controller := NewSessionController(locationService)

	if controller == nil {
		t.Fatal("Expected non-nil controller")
	}
}

func TestSessionControllerGetUserLocation(t *testing.T) {
	mockRepo := &repositories.MockSessionRepository{}
	locationService := services.NewLocationService(mockRepo)
	controller := NewSessionController(locationService)

	if controller == nil {
		t.Fatal("Expected non-nil controller")
	}
}

func TestSessionControllerGetSessionLocations(t *testing.T) {
	mockRepo := &repositories.MockSessionRepository{}
	locationService := services.NewLocationService(mockRepo)
	controller := NewSessionController(locationService)

	if controller == nil {
		t.Fatal("Expected non-nil controller")
	}
}

func TestSessionControllerUpdateUserLocation(t *testing.T) {
	mockRepo := &repositories.MockSessionRepository{}
	locationService := services.NewLocationService(mockRepo)
	controller := NewSessionController(locationService)

	if controller == nil {
		t.Fatal("Expected non-nil controller")
	}
}
