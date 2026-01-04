package config

import (
	"os"
	"testing"
)

func TestNewMockConfig(t *testing.T) {
	cfg := NewMockConfig()

	if cfg == nil {
		t.Fatal("Expected non-nil Config")
	}

	if cfg.Server.Port != 8080 {
		t.Fatalf("Expected port 8080, got %d", cfg.Server.Port)
	}

	if cfg.Server.Env != "development" {
		t.Fatalf("Expected environment 'development', got %s", cfg.Server.Env)
	}

	if cfg.Database.Type != "sqlite" {
		t.Fatalf("Expected database type 'sqlite', got %s", cfg.Database.Type)
	}

	if cfg.Database.Connection != ":memory:" {
		t.Fatalf("Expected database connection ':memory:', got %s", cfg.Database.Connection)
	}
}

func TestNewMockConfigWithValues(t *testing.T) {
	cfg := NewMockConfigWithValues(3000, "staging", "mysql", "user:pass@tcp(host:3306)/db")

	if cfg.Server.Port != 3000 {
		t.Fatalf("Expected port 3000, got %d", cfg.Server.Port)
	}

	if cfg.Server.Env != "staging" {
		t.Fatalf("Expected environment 'staging', got %s", cfg.Server.Env)
	}

	if cfg.Database.Type != "mysql" {
		t.Fatalf("Expected database type 'mysql', got %s", cfg.Database.Type)
	}

	if cfg.Database.Connection != "user:pass@tcp(host:3306)/db" {
		t.Fatalf("Expected database connection 'user:pass@tcp(host:3306)/db', got %s", cfg.Database.Connection)
	}
}

func TestDefaultConfigLoader_LoadDevelopment(t *testing.T) {
	os.Setenv("ENVIRONMENT", "dev")
	defer os.Unsetenv("ENVIRONMENT")

	loader := &DefaultConfigLoader{}
	cfg := loader.LoadConfig()

	// Should load from .env.dev file
	if cfg == nil {
		t.Fatal("Expected non-nil Config")
	}

	if cfg.Server.Env != "dev" {
		t.Fatalf("Expected environment 'dev', got %s", cfg.Server.Env)
	}

	if cfg.Server.Port != 8080 {
		t.Fatalf("Expected port 8080 from .env, got %d", cfg.Server.Port)
	}

	if cfg.Database.Type != "sqlite" {
		t.Fatalf("Expected database type 'sqlite' from .env, got %s", cfg.Database.Type)
	}

	if cfg.Database.Connection != ":memory:" {
		t.Fatalf("Expected database connection ':memory:' from .env, got %s", cfg.Database.Connection)
	}
}

func TestDefaultConfigLoader_EmptyEnvironmentDefaultsToDevelopment(t *testing.T) {
	// Clear ENVIRONMENT variable
	os.Unsetenv("ENVIRONMENT")
	defer os.Setenv("ENVIRONMENT", os.Getenv("ENVIRONMENT"))

	loader := &DefaultConfigLoader{}
	cfg := loader.LoadConfig()

	if cfg.Server.Env != "dev" {
		t.Fatalf("Expected environment 'dev', got %s", cfg.Server.Env)
	}
}

func TestDefaultConfigLoader_MissingPort(t *testing.T) {
	os.Setenv("ENVIRONMENT", "dev")
	defer os.Unsetenv("ENVIRONMENT")

	loader := &DefaultConfigLoader{}
	cfg := loader.LoadConfig()

	// Should load PORT from .env.dev file successfully
	if cfg.Server.Port != 8080 {
		t.Fatalf("Expected port 8080 from .env file, got %d", cfg.Server.Port)
	}
}

func TestDefaultConfigLoader_PortFromFile(t *testing.T) {
	os.Setenv("ENVIRONMENT", "dev")
	defer os.Unsetenv("ENVIRONMENT")

	loader := &DefaultConfigLoader{}
	cfg := loader.LoadConfig()

	// Should load PORT from .env.dev file
	if cfg.Server.Port != 8080 {
		t.Fatalf("Expected port 8080 from .env file, got %d", cfg.Server.Port)
	}
}

func TestDefaultConfigLoader_DatabaseTypeFromFile(t *testing.T) {
	os.Setenv("ENVIRONMENT", "dev")
	defer os.Unsetenv("ENVIRONMENT")

	loader := &DefaultConfigLoader{}
	cfg := loader.LoadConfig()

	// Should load DB_TYPE from .env.dev file
	if cfg.Database.Type != "sqlite" {
		t.Fatalf("Expected database type 'sqlite' from .env file, got %s", cfg.Database.Type)
	}
}

func TestDefaultConfigLoader_DatabaseConnectionFromFile(t *testing.T) {
	os.Setenv("ENVIRONMENT", "dev")
	defer os.Unsetenv("ENVIRONMENT")

	loader := &DefaultConfigLoader{}
	cfg := loader.LoadConfig()

	// Should load DB_CONNECTION from .env.dev file
	if cfg.Database.Connection != ":memory:" {
		t.Fatalf("Expected connection ':memory:' from .env file, got %s", cfg.Database.Connection)
	}
}
