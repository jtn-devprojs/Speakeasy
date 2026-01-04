package config

// ConfigLoader defines the interface for loading configuration
type ConfigLoader interface {
	Load(environment string) (*Config, error)
}
