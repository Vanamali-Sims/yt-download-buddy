package config

import (
	"os"
)

// Config structure to hold application configuration
type Config struct {
	OutputPath  string
	VideoFormat string
}

// LoadConfig loads configuration from environment variables or defaults
func LoadConfig() *Config {
	return &Config{
		OutputPath:  getEnv("OUTPUT_PATH", "./downloads/"),
		VideoFormat: getEnv("VIDEO_FORMAT", "best"),
	}
}

// Helper function to read environment variables or fallback to default value
func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}
