package config

import (
	"os"
)

type Config struct {
	DatabaseURL string
	ServerPort  string
}

func LoadConfig() *Config {
	dbURL := os.Getenv("DB_URL")
	if dbURL == "" {
		// Дефолтная строка для локального запуска без Docker
		dbURL = "postgres://postgres:mari1205@localhost:5432/theater_db?sslmode=disable"
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = ":8080"
	}

	return &Config{
		DatabaseURL: dbURL,
		ServerPort:  port,
	}
}
