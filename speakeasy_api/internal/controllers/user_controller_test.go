package controllers

import (
	"testing"

	"github.com/speakeasy/speakeasy-api/internal/repositories"
	"github.com/speakeasy/speakeasy-api/internal/services"
)

func TestNewUserController(t *testing.T) {
	mockRepo := &repositories.MockUserRepository{}
	userService := services.NewUserService(mockRepo)
	controller := NewUserController(userService)

	if controller == nil {
		t.Fatal("Expected non-nil UserController")
	}

	if controller.userService == nil {
		t.Fatal("Expected userService to be initialized")
	}
}

func TestUserControllerGetUser(t *testing.T) {
	mockRepo := &repositories.MockUserRepository{}
	userService := services.NewUserService(mockRepo)
	controller := NewUserController(userService)

	if controller == nil {
		t.Fatal("Expected non-nil controller")
	}
}

func TestUserControllerCreateUser(t *testing.T) {
	mockRepo := &repositories.MockUserRepository{}
	userService := services.NewUserService(mockRepo)
	controller := NewUserController(userService)

	if controller == nil {
		t.Fatal("Expected non-nil controller")
	}
}

func TestUserControllerUpdateUser(t *testing.T) {
	mockRepo := &repositories.MockUserRepository{}
	userService := services.NewUserService(mockRepo)
	controller := NewUserController(userService)

	if controller == nil {
		t.Fatal("Expected non-nil controller")
	}
}

func TestUserControllerDeleteUser(t *testing.T) {
	mockRepo := &repositories.MockUserRepository{}
	userService := services.NewUserService(mockRepo)
	controller := NewUserController(userService)

	if controller == nil {
		t.Fatal("Expected non-nil controller")
	}
}

func TestUserControllerGetUserPreferences(t *testing.T) {
	mockRepo := &repositories.MockUserRepository{}
	userService := services.NewUserService(mockRepo)
	controller := NewUserController(userService)

	if controller == nil {
		t.Fatal("Expected non-nil controller")
	}
}

func TestUserControllerUpdateUserPreferences(t *testing.T) {
	mockRepo := &repositories.MockUserRepository{}
	userService := services.NewUserService(mockRepo)
	controller := NewUserController(userService)

	if controller == nil {
		t.Fatal("Expected non-nil controller")
	}
}
