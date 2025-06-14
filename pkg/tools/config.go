package tools

import "os"

type DatabaseConfig struct {
	ConnStr string
}

type ServerConfig struct {
	Port string
	Host string
}

type Config struct {
	Postgres DatabaseConfig
	Server   ServerConfig
}

func New() *Config {
	return &Config{
		Postgres: DatabaseConfig{
			ConnStr: os.Getenv("DB_ADDRESS"),
		},
		Server: ServerConfig{
			Port: os.Getenv("SERVER_PORT"),
			Host: os.Getenv("SERVER_HOST"),
		},
	}
}
