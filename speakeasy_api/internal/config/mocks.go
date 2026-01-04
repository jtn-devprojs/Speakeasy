package config

func NewMockConfig() *Config {
	return &Config{
		Server: ServerConfig{
			Port: 8080,
			Env:  "development",
		},
		Database: DatabaseConfig{
			Type:       "sqlite",
			Connection: ":memory:",
		},
	}
}

func NewMockConfigWithValues(port int, env, dbType, dbConnection string) *Config {
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

type MockConfigLoader struct {
	LoadFunc func(environment string) (*Config, error)
}

func (mcl *MockConfigLoader) Load(environment string) (*Config, error) {
	if mcl.LoadFunc != nil {
		return mcl.LoadFunc(environment)
	}
	return NewMockConfig(), nil
}

func NewMockConfigLoader() *MockConfigLoader {
	return &MockConfigLoader{
		LoadFunc: func(environment string) (*Config, error) {
			return NewMockConfig(), nil
		},
	}
}
