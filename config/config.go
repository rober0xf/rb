package config

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
)

type Config struct {
	APIKey string
}

func Load() (*Config, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return nil, fmt.Errorf("failed to get home dir: %w", err)
	}

	envPath := filepath.Join(homeDir, ".rb", ".env")
	_ = godotenv.Load(envPath)

	apiKey := getEnv("API_KEY", "")
	if apiKey == "" {
		return nil, fmt.Errorf("API_KEY is not configured")
	}

	return &Config{
		APIKey: apiKey,
	}, nil
}

func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	return fallback
}
