package config

import (
	"errors"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	AppName string
	Port    string
}

func LoadConfig() (*Config, error) {
	err := godotenv.Load()
	if err != nil {
		return nil, errors.New("failed to load .env file")
	}

	config := &Config{
		AppName: os.Getenv("APP_NAME"),
		Port:    os.Getenv("PORT"),
	}

	if config.AppName == "" || config.Port == "" {
		return nil, errors.New("required environment variables missing")
	}

	return config, nil
}
