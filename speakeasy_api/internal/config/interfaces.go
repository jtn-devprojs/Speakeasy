package config

// ConfigLoader defines the interface for loading configuration
type ConfigLoader interface {
	LoadConfig() *Config
	Load(environment string) *Config
	LoadEnvFile(environment string) (map[string]string, error)
}
