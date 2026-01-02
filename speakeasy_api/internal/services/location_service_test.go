package services

import (
	"testing"

	"github.com/speakeasy/speakeasy-api/internal/repositories"
)

func TestNewLocationService(t *testing.T) {
	mockRepo := &repositories.MockSessionRepository{}
	service := NewLocationService(mockRepo)

	if service == nil {
		t.Fatal("Expected non-nil LocationService")
	}
}

func TestLocationServiceIsUserInVicinity(t *testing.T) {
	mockRepo := &repositories.MockSessionRepository{}
	service := NewLocationService(mockRepo)

	result := service.IsUserInVicinity(40.7128, -74.0060, 40.7128, -74.0060, 1.0)
	if result != false {
		t.Fatal("Expected false for test vicinity check")
	}
}

func TestLocationServiceGetNearbyLocations(t *testing.T) {
	mockRepo := &repositories.MockSessionRepository{}
	service := NewLocationService(mockRepo)

	result, err := service.GetNearbyLocations(40.7128, -74.0060, 5.0)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	if result != nil {
		t.Fatal("Expected nil result for unimplemented method")
	}
}

func TestLocationServiceGetUserLocation(t *testing.T) {
	mockRepo := &repositories.MockSessionRepository{}
	service := NewLocationService(mockRepo)

	result, err := service.GetUserLocation("user123")
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	if result != nil {
		t.Fatal("Expected nil result for unimplemented method")
	}
}

func TestLocationServiceUpdateUserLocation(t *testing.T) {
	mockRepo := &repositories.MockSessionRepository{}
	service := NewLocationService(mockRepo)

	err := service.UpdateUserLocation("user123", 40.7128, -74.0060)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
}
