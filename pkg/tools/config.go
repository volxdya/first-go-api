package tools

import "os"

type DatabaseConfig struct {
	ConnStr string
}

type Config struct {
	Postgres DatabaseConfig
}

func New() *Config {
	return &Config{
		Postgres: DatabaseConfig{
			ConnStr: os.Getenv("DB_ADDRESS"),
		},
	}
}
