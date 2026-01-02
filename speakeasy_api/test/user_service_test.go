package main

import (
	"testing"

	"github.com/speakeasy/speakeasy-api/internal/services"
)

func TestNewUserService(t *testing.T) {
	service := services.NewUserService()

	if service == nil {
		t.Fatal("Expected non-nil UserService")
	}
}

func TestUserServiceGetUserByID(t *testing.T) {
	service := services.NewUserService()

	_, err := service.GetUserByID("user123")
	if err != services.ErrNotImplemented {
		t.Fatalf("Expected ErrNotImplemented, got %v", err)
	}
}

func TestUserServiceCreateUser(t *testing.T) {
	service := services.NewUserService()

	_, err := service.CreateUser("testuser", "test@example.com", "password")
	if err != services.ErrNotImplemented {
		t.Fatalf("Expected ErrNotImplemented, got %v", err)
	}
}

func TestUserServiceUpdateUser(t *testing.T) {
	service := services.NewUserService()

	err := service.UpdateUser("user123", map[string]interface{}{})
	if err != services.ErrNotImplemented {
		t.Fatalf("Expected ErrNotImplemented, got %v", err)
	}
}

func TestUserServiceDeleteUser(t *testing.T) {
	service := services.NewUserService()

	err := service.DeleteUser("user123")
	if err != services.ErrNotImplemented {
		t.Fatalf("Expected ErrNotImplemented, got %v", err)
	}
}

func TestUserServiceGetUserPreferences(t *testing.T) {
	service := services.NewUserService()

	_, err := service.GetUserPreferences("user123")
	if err != services.ErrNotImplemented {
		t.Fatalf("Expected ErrNotImplemented, got %v", err)
	}
}

func TestUserServiceUpdateUserPreferences(t *testing.T) {
	service := services.NewUserService()

	err := service.UpdateUserPreferences("user123", map[string]interface{}{})
	if err != services.ErrNotImplemented {
		t.Fatalf("Expected ErrNotImplemented, got %v", err)
	}
}
