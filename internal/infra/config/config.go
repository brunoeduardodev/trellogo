package config

import "os"

type DatabaseConfig struct {
	Port     string
	Password string
	User     string
	Name     string
	Host     string
}

type RestApiConfig struct {
	Port string
}

type Config struct {
	RestApi  RestApiConfig
	Database DatabaseConfig
}

func NewConfig() *Config {
	return &Config{
		Database: DatabaseConfig{
			User:     os.Getenv("DB_USER"),
			Port:     os.Getenv("DB_PORT"),
			Password: os.Getenv("DB_PASSWORD"),
			Name:     os.Getenv("DB_NAME"),
			Host:     os.Getenv("DB_HOST"),
		},

		RestApi: RestApiConfig{
			Port: os.Getenv("PORT"),
		},
	}
}
