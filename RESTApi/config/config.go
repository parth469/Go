package config

import (
	"errors"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	AppName  string
	Port     string
	Database string
}

func LoadConfig() (*Config, error) {
	err := godotenv.Load()
	if err != nil {
		return nil, errors.New("failed to load .env file")
	}

	config := &Config{
		AppName:  os.Getenv("APP_NAME"),
		Port:     os.Getenv("PORT"),
		Database: os.Getenv("DATABASE"),
	}

	if config.AppName == "" || config.Port == "" || config.Database == "" {
		return nil, errors.New("required environment variables missing")
	}

	return config, nil
}
