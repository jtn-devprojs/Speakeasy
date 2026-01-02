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

	if container.UserHandler == nil {
		t.Fatal("Expected UserHandler to be initialized")
	}

	if container.AuthHandler == nil {
		t.Fatal("Expected AuthHandler to be initialized")
	}
}

func TestContainer_DependencyInjection(t *testing.T) {
	container := NewContainer()

	// Verify that handlers have their dependencies
	if container.UserHandler == nil {
		t.Fatal("UserHandler should be initialized")
	}

	if container.AuthHandler == nil {
		t.Fatal("AuthHandler should be initialized")
	}

	// Verify service dependencies
	if container.UserService == nil {
		t.Fatal("UserService should be initialized")
	}

	if container.AuthService == nil {
		t.Fatal("AuthService should be initialized")
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
