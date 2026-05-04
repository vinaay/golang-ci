package config

import (
	"os"
	"testing"
	"time"
)

func TestConfig_Validate(t *testing.T) {
	tests := []struct {
		name    string
		config  *Config
		wantErr bool
	}{
		{
			name: "valid config",
			config: &Config{
				Port:         "3000",
				DataFilePath: "./data.json",
				LogLevel:     "info",
				JSONLog:      false,
				ReadTimeout:  10 * time.Second,
				WriteTimeout: 10 * time.Second,
			},
			wantErr: false,
		},
		{
			name: "empty port",
			config: &Config{
				Port:         "",
				DataFilePath: "./data.json",
				LogLevel:     "info",
			},
			wantErr: true,
		},
		{
			name: "empty data file path",
			config: &Config{
				Port:         "3000",
				DataFilePath: "",
				LogLevel:     "info",
			},
			wantErr: true,
		},
		{
			name: "invalid log level",
			config: &Config{
				Port:         "3000",
				DataFilePath: "./data.json",
				LogLevel:     "invalid",
			},
			wantErr: true,
		},
		{
			name: "valid log levels",
			config: &Config{
				Port:         "3000",
				DataFilePath: "./data.json",
				LogLevel:     "debug",
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.config.Validate()
			if (err != nil) != tt.wantErr {
				t.Errorf("Config.Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestLoad(t *testing.T) {
	// Save original env values
	originalPort := os.Getenv("PORT")
	originalLogLevel := os.Getenv("LOG_LEVEL")

	// Clean up
	defer func() {
		if originalPort != "" {
			os.Setenv("PORT", originalPort)
		} else {
			os.Unsetenv("PORT")
		}
		if originalLogLevel != "" {
			os.Setenv("LOG_LEVEL", originalLogLevel)
		} else {
			os.Unsetenv("LOG_LEVEL")
		}
	}()

	// Test with defaults
	os.Unsetenv("PORT")
	os.Unsetenv("LOG_LEVEL")
	cfg, err := Load()
	if err != nil {
		t.Fatalf("Load() error = %v", err)
	}
	if cfg.Port != "3000" {
		t.Errorf("Load() Port = %v, want %v", cfg.Port, "3000")
	}
	if cfg.LogLevel != "info" {
		t.Errorf("Load() LogLevel = %v, want %v", cfg.LogLevel, "info")
	}

	// Test with environment variables
	os.Setenv("PORT", "8080")
	os.Setenv("LOG_LEVEL", "debug")
	cfg, err = Load()
	if err != nil {
		t.Fatalf("Load() error = %v", err)
	}
	if cfg.Port != "8080" {
		t.Errorf("Load() Port = %v, want %v", cfg.Port, "8080")
	}
	if cfg.LogLevel != "debug" {
		t.Errorf("Load() LogLevel = %v, want %v", cfg.LogLevel, "debug")
	}
}
