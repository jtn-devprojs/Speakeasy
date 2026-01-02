package config

import (
	"os"
	"strconv"
)

type ServerConfig struct {
	Port int
	Env  string
}

type DatabaseConfig struct {
	Type       string
	Connection string
}

type Config struct {
	Server   ServerConfig
	Database DatabaseConfig
}

func LoadConfig() *Config {
	env := os.Getenv("ENVIRONMENT")
	if env == "" {
		env = "development"
	}

	port := 8080
	if portStr := os.Getenv("PORT"); portStr != "" {
		if p, err := strconv.Atoi(portStr); err == nil {
			port = p
		}
	}

	dbType := os.Getenv("DB_TYPE")
	if dbType == "" {
		dbType = getDefaultDBType(env)
	}

	dbConnection := os.Getenv("DB_CONNECTION")
	if dbConnection == "" {
		dbConnection = getDefaultConnection(env, dbType)
	}

	return &Config{
		Server: ServerConfig{
			Port: port,
			Env:  env,
		},
		Database: DatabaseConfig{
			Type:       dbType,
			Connection: dbConnection,
		},
	}
}

func getDefaultDBType(env string) string {
	switch env {
	case "development":
		return "sqlite"
	case "staging", "production":
		return "mysql"
	default:
		return "sqlite"
	}
}

func getDefaultConnection(env, dbType string) string {
	switch env {
	case "development":
		if dbType == "sqlite" {
			return ":memory:"
		}
	case "staging":
		if dbType == "mysql" {
			return "user:password@tcp(staging-db:3306)/speakeasy"
		}
	case "production":
		if dbType == "mysql" {
			return "user:password@tcp(prod-db:3306)/speakeasy"
		}
	}
	return ":memory:"
}
