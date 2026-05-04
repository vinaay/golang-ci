// Package logger provides structured logging functionality using Go's log/slog package.
package logger

import (
	"log/slog"
	"os"
)

var (
	// Logger is the global structured logger instance.
	Logger *slog.Logger
)

// Init initializes the global logger with the specified log level.
// If jsonOutput is true, logs will be in JSON format.
func Init(level string, jsonOutput bool) {
	var logLevel slog.Level
	switch level {
	case "debug":
		logLevel = slog.LevelDebug
	case "info":
		logLevel = slog.LevelInfo
	case "warn":
		logLevel = slog.LevelWarn
	case "error":
		logLevel = slog.LevelError
	default:
		logLevel = slog.LevelInfo
	}

	opts := &slog.HandlerOptions{
		Level: logLevel,
	}

	var handler slog.Handler
	if jsonOutput {
		handler = slog.NewJSONHandler(os.Stdout, opts)
	} else {
		handler = slog.NewTextHandler(os.Stdout, opts)
	}

	Logger = slog.New(handler)
}

// Debug logs a debug message.
func Debug(msg string, args ...any) {
	if Logger != nil {
		Logger.Debug(msg, args...)
	}
}

// Info logs an info message.
func Info(msg string, args ...any) {
	if Logger != nil {
		Logger.Info(msg, args...)
	}
}

// Warn logs a warning message.
func Warn(msg string, args ...any) {
	if Logger != nil {
		Logger.Warn(msg, args...)
	}
}

// Error logs an error message.
func Error(msg string, args ...any) {
	if Logger != nil {
		Logger.Error(msg, args...)
	}
}
