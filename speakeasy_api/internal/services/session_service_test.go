package services

import (
	"testing"

	"github.com/speakeasy/speakeasy-api/internal/repositories"
)

func TestNewSessionService(t *testing.T) {
	mockSessionRepo := &repositories.MockSessionRepository{}
	mockSessionUserRepo := &repositories.MockSessionUserRepository{}
	service := NewSessionService(mockSessionRepo, mockSessionUserRepo)

	if service == nil {
		t.Fatal("Expected non-nil SessionService")
	}
}

func TestSessionService_JoinSession(t *testing.T) {
	mockSessionRepo := &repositories.MockSessionRepository{}
	mockSessionUserRepo := &repositories.MockSessionUserRepository{}
	service := NewSessionService(mockSessionRepo, mockSessionUserRepo)

	// Configure mock to succeed
	mockSessionUserRepo.JoinSessionWithLockFunc = func(sessionID, userID string) error {
		return nil
	}

	err := service.JoinSession("session-1", "user-1")
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	// Verify the repository method was called with correct parameters
	if mockSessionUserRepo.JoinSessionWithLockFunc == nil {
		t.Fatal("Expected JoinSessionWithLock to be called")
	}
}

func TestSessionService_JoinSession_Error(t *testing.T) {
	mockSessionRepo := &repositories.MockSessionRepository{}
	mockSessionUserRepo := &repositories.MockSessionUserRepository{}
	service := NewSessionService(mockSessionRepo, mockSessionUserRepo)

	// Configure mock to fail
	mockSessionUserRepo.JoinSessionWithLockFunc = func(sessionID, userID string) error {
		return &repositories.MockError{Message: "session not found"}
	}

	err := service.JoinSession("non-existent", "user-1")
	if err == nil {
		t.Fatal("Expected error when session not found")
	}
}

func TestSessionService_LeaveSession(t *testing.T) {
	mockSessionRepo := &repositories.MockSessionRepository{}
	mockSessionUserRepo := &repositories.MockSessionUserRepository{}
	service := NewSessionService(mockSessionRepo, mockSessionUserRepo)

	// Configure mock to succeed
	mockSessionUserRepo.LeaveSessionWithCleanupFunc = func(sessionID, userID string) error {
		return nil
	}

	err := service.LeaveSession("session-1", "user-1")
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	// Verify the repository method was called
	if mockSessionUserRepo.LeaveSessionWithCleanupFunc == nil {
		t.Fatal("Expected LeaveSessionWithCleanup to be called")
	}
}

func TestSessionService_LeaveSession_Error(t *testing.T) {
	mockSessionRepo := &repositories.MockSessionRepository{}
	mockSessionUserRepo := &repositories.MockSessionUserRepository{}
	service := NewSessionService(mockSessionRepo, mockSessionUserRepo)

	// Configure mock to fail
	mockSessionUserRepo.LeaveSessionWithCleanupFunc = func(sessionID, userID string) error {
		return &repositories.MockError{Message: "user not in session"}
	}

	err := service.LeaveSession("session-1", "unknown-user")
	if err == nil {
		t.Fatal("Expected error when user not in session")
	}
}

func TestSessionServiceIsUserInVicinity(t *testing.T) {
	mockSessionRepo := &repositories.MockSessionRepository{}
	mockSessionUserRepo := &repositories.MockSessionUserRepository{}
	service := NewSessionService(mockSessionRepo, mockSessionUserRepo)

	result := service.IsUserInVicinity(40.7128, -74.0060, 40.7128, -74.0060, 1.0)
	if result != false {
		t.Fatal("Expected false for test vicinity check")
	}
}

func TestSessionServiceGetNearbyLocations(t *testing.T) {
	mockSessionRepo := &repositories.MockSessionRepository{}
	mockSessionUserRepo := &repositories.MockSessionUserRepository{}
	service := NewSessionService(mockSessionRepo, mockSessionUserRepo)

	result, err := service.GetNearbyLocations(40.7128, -74.0060, 5.0)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	if result != nil {
		t.Fatal("Expected nil result for unimplemented method")
	}
}

func TestSessionServiceGetUserLocation(t *testing.T) {
	mockSessionRepo := &repositories.MockSessionRepository{}
	mockSessionUserRepo := &repositories.MockSessionUserRepository{}
	service := NewSessionService(mockSessionRepo, mockSessionUserRepo)

	result, err := service.GetUserLocation("user123")
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	if result != nil {
		t.Fatal("Expected nil result for unimplemented method")
	}
}

func TestSessionServiceUpdateUserLocation(t *testing.T) {
	mockSessionRepo := &repositories.MockSessionRepository{}
	mockSessionUserRepo := &repositories.MockSessionUserRepository{}
	service := NewSessionService(mockSessionRepo, mockSessionUserRepo)

	err := service.UpdateUserLocation("user123", 40.7128, -74.0060)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
}
