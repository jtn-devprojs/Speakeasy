package config

import (
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
