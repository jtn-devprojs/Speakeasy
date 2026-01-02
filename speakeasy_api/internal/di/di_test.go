package di

import (
	"testing"
)

func TestNewContainer(t *testing.T) {
	container := NewContainer()

	if container == nil {
		t.Fatal("Expected non-nil container")
	}

	if container.UserService == nil {
		t.Fatal("Expected UserService to be initialized")
	}

	if container.AuthService == nil {
		t.Fatal("Expected AuthService to be initialized")
	}

	if container.LocationService == nil {
		t.Fatal("Expected LocationService to be initialized")
	}

	if container.UserController == nil {
		t.Fatal("Expected UserController to be initialized")
	}

	if container.AuthController == nil {
		t.Fatal("Expected AuthController to be initialized")
	}

	if container.LocationController == nil {
		t.Fatal("Expected LocationController to be initialized")
	}
}

func TestContainer_DependencyInjection(t *testing.T) {
	container := NewContainer()

	// Verify that controllers have their dependencies
	if container.UserController == nil {
		t.Fatal("UserController should be initialized")
	}

	if container.AuthController == nil {
		t.Fatal("AuthController should be initialized")
	}

	if container.LocationController == nil {
		t.Fatal("LocationController should be initialized")
	}

	// Verify service dependencies
	if container.UserService == nil {
		t.Fatal("UserService should be initialized")
	}

	if container.AuthService == nil {
		t.Fatal("AuthService should be initialized")
	}

	if container.LocationService == nil {
		t.Fatal("LocationService should be initialized")
	}
}

func TestContainer_Singleton(t *testing.T) {
	container1 := NewContainer()
	container2 := NewContainer()

	// Containers are separate instances (not singletons)
	if container1 == container2 {
		t.Fatal("Expected separate container instances")
	}

	// But services within a container are consistent
	if container1.UserService != container1.UserService {
		t.Fatal("UserService should be consistent within container")
	}
}
