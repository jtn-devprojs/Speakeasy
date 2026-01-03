package controllers

import (
	"testing"

	"github.com/speakeasy/speakeasy-api/internal/repositories"
	"github.com/speakeasy/speakeasy-api/internal/services"
)

func TestNewSessionController(t *testing.T) {
	mockRepo := &repositories.MockSessionRepository{}
	sessionService := services.NewSessionService(mockRepo)
	controller := NewSessionController(sessionService)

	if controller == nil {
		t.Fatal("Expected non-nil SessionController")
	}

	if controller.sessionService == nil {
		t.Fatal("Expected sessionService to be initialized")
	}
}

func TestSessionControllerCheckVicinity(t *testing.T) {
	mockRepo := &repositories.MockSessionRepository{}
	sessionService := services.NewSessionService(mockRepo)
	controller := NewSessionController(sessionService)

	if controller == nil {
		t.Fatal("Expected non-nil controller")
	}
}

func TestSessionControllerGetClosestSession(t *testing.T) {
	mockRepo := &repositories.MockSessionRepository{}
	sessionService := services.NewSessionService(mockRepo)
	controller := NewSessionController(sessionService)

	if controller == nil {
		t.Fatal("Expected non-nil controller")
	}
}

func TestSessionControllerGetNearbyLocations(t *testing.T) {
	mockRepo := &repositories.MockSessionRepository{}
	sessionService := services.NewSessionService(mockRepo)
	controller := NewSessionController(sessionService)

	if controller == nil {
		t.Fatal("Expected non-nil controller")
	}
}

func TestSessionControllerGetUserLocation(t *testing.T) {
	mockRepo := &repositories.MockSessionRepository{}
	sessionService := services.NewSessionService(mockRepo)
	controller := NewSessionController(sessionService)

	if controller == nil {
		t.Fatal("Expected non-nil controller")
	}
}

func TestSessionControllerGetSessionLocations(t *testing.T) {
	mockRepo := &repositories.MockSessionRepository{}
	sessionService := services.NewSessionService(mockRepo)
	controller := NewSessionController(sessionService)

	if controller == nil {
		t.Fatal("Expected non-nil controller")
	}
}

func TestSessionControllerUpdateUserLocation(t *testing.T) {
	mockRepo := &repositories.MockSessionRepository{}
	sessionService := services.NewSessionService(mockRepo)
	controller := NewSessionController(sessionService)

	if controller == nil {
		t.Fatal("Expected non-nil controller")
	}
}
