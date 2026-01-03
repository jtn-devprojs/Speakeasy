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
