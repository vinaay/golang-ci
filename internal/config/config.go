// Package config provides configuration management for the application.
package config

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
)

// Config holds the application configuration.
type Config struct {
	Port         string
	DataFilePath string
	LogLevel     string
	JSONLog      bool
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

// Load loads configuration from environment variables with defaults.
// It also attempts to load a .env file if present.
func Load() (*Config, error) {
	// Try to load .env file, but don't fail if it doesn't exist
	_ = godotenv.Load()

	cfg := &Config{
		Port:         getEnv("PORT", "3000"),
		DataFilePath: getEnv("DATA_FILE_PATH", "./data.json"),
		LogLevel:     getEnv("LOG_LEVEL", "info"),
		JSONLog:      getEnvBool("JSON_LOG", false),
		ReadTimeout:  getEnvDuration("READ_TIMEOUT", 10*time.Second),
		WriteTimeout: getEnvDuration("WRITE_TIMEOUT", 10*time.Second),
	}

	if err := cfg.Validate(); err != nil {
		return nil, fmt.Errorf("invalid configuration: %w", err)
	}

	return cfg, nil
}

// Validate validates the configuration values.
func (c *Config) Validate() error {
	if c.Port == "" {
		return fmt.Errorf("port cannot be empty")
	}

	if c.DataFilePath == "" {
		return fmt.Errorf("data file path cannot be empty")
	}

	// Validate log level
	validLogLevels := map[string]bool{
		"debug": true,
		"info":  true,
		"warn":  true,
		"error": true,
	}
	if !validLogLevels[c.LogLevel] {
		return fmt.Errorf("invalid log level: %s (must be debug, info, warn, or error)", c.LogLevel)
	}

	return nil
}

// getEnv retrieves an environment variable or returns a default value.
func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

// getEnvBool retrieves a boolean environment variable or returns a default value.
func getEnvBool(key string, defaultValue bool) bool {
	if value := os.Getenv(key); value != "" {
		if parsed, err := strconv.ParseBool(value); err == nil {
			return parsed
		}
	}
	return defaultValue
}

// getEnvDuration retrieves a duration environment variable or returns a default value.
func getEnvDuration(key string, defaultValue time.Duration) time.Duration {
	if value := os.Getenv(key); value != "" {
		if parsed, err := time.ParseDuration(value); err == nil {
			return parsed
		}
	}
	return defaultValue
}
