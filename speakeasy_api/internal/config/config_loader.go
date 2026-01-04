package config

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
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

type DefaultConfigLoader struct{}

func (dcl *DefaultConfigLoader) LoadConfig() *Config {
	env := os.Getenv("ENVIRONMENT")
	if env == "" {
		env = "dev"
	}

	return dcl.Load(env)
}

// Load loads configuration for a specific environment
func (dcl *DefaultConfigLoader) Load(environment string) *Config {

	envSettings, err := dcl.LoadEnvFile(environment)
	if err != nil {
		panic(fmt.Sprintf("failed to load settings file: %v", err))
	}

	// Set all settings as environment variables
	for key, value := range envSettings {
		os.Setenv(key, value)
	}

	return &Config{
		Server: ServerConfig{
			Port: dcl.loadAndValidateInteger("PORT"),
			Env:  environment,
		},
		Database: DatabaseConfig{
			Type:       dcl.loadAndValidateConfig("DB_TYPE"),
			Connection: dcl.loadAndValidateConfig("DB_CONNECTION"),
		},
	}
}

// loadAndValidateInteger retrieves and validates an integer environment variable
func (dcl *DefaultConfigLoader) loadAndValidateInteger(key string) int {
	portStr := dcl.loadAndValidateConfig(key)
	port, err := strconv.Atoi(portStr)
	if err != nil {
		panic(fmt.Sprintf("invalid PORT value: %s", portStr))
	}
	return port
}

// loadAndValidateConfig retrieves an environment variable and panics if it's not found
func (dcl *DefaultConfigLoader) loadAndValidateConfig(key string) string {
	value := os.Getenv(key)
	if value == "" {
		panic(fmt.Sprintf("%s not found in settings file", key))
	}
	return value
}

// TODO: Add a configuration file validator to check for missing or invalid settings

// LoadConfig is the package-level convenience function to load configuration
// using the default ConfigLoader implementation
func LoadConfig() *Config {
	loader := &DefaultConfigLoader{}
	return loader.LoadConfig()
}

// LoadEnvFile reads a .env file and returns a map of key-value pairs
func (dcl *DefaultConfigLoader) LoadEnvFile(environment string) (map[string]string, error) {
	settings := make(map[string]string)

	envFile := fmt.Sprintf(".env.%s", environment)

	// Check if file exists
	if info, err := os.Stat(envFile); err != nil || info.IsDir() {
		return nil, fmt.Errorf("settings file not found: %s", envFile)
	}

	file, err := os.Open(envFile)
	if err != nil {
		return nil, fmt.Errorf("failed to open settings file: %w", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		// Skip empty lines and comments
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}

		// Parse KEY=VALUE
		parts := strings.SplitN(line, "=", 2)
		if len(parts) == 2 {
			key := strings.TrimSpace(parts[0])
			value := strings.TrimSpace(parts[1])
			settings[key] = value
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("error reading settings file: %w", err)
	}

	return settings, nil
}
