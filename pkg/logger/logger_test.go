package logger

import (
	"testing"
)

func TestInit(t *testing.T) {
	// Test text output
	Init("info", false)
	if Logger == nil {
		t.Error("Logger should not be nil after Init")
	}

	// Test JSON output
	Init("debug", true)
	if Logger == nil {
		t.Error("Logger should not be nil after Init")
	}

	// Test invalid log level (should default to info)
	Init("invalid", false)
	if Logger == nil {
		t.Error("Logger should not be nil after Init")
	}
}

func TestLogFunctions(_ *testing.T) {
	Init("debug", false)

	// These should not panic
	Debug("test debug message", "key", "value")
	Info("test info message", "key", "value")
	Warn("test warn message", "key", "value")
	Error("test error message", "key", "value")

	// Test with nil logger (should not panic)
	Logger = nil
	Debug("test", "key", "value")
	Info("test", "key", "value")
	Warn("test", "key", "value")
	Error("test", "key", "value")
}
