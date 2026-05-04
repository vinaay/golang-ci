// Package middleware provides HTTP middleware functions for request processing.
package middleware

import (
	"time"

	"api/pkg/logger"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

const (
	// RequestIDHeader is the HTTP header key for request ID.
	RequestIDHeader = "X-Request-ID"
	// RequestIDKey is the context key for request ID.
	RequestIDKey = "request_id"
)

// RequestID adds a unique request ID to each request.
func RequestID() gin.HandlerFunc {
	return func(c *gin.Context) {
		requestID := c.GetHeader(RequestIDHeader)
		if requestID == "" {
			requestID = uuid.New().String()
		}
		c.Set(RequestIDKey, requestID)
		c.Header(RequestIDHeader, requestID)
		c.Next()
	}
}

// Logger logs HTTP requests and responses.
func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		path := c.Request.URL.Path
		raw := c.Request.URL.RawQuery

		// Process request
		c.Next()

		// Calculate latency
		latency := time.Since(start)

		// Get request ID
		requestID, _ := c.Get(RequestIDKey)
		requestIDStr := ""
		if id, ok := requestID.(string); ok {
			requestIDStr = id
		}

		// Log request
		logger.Info("HTTP Request",
			"method", c.Request.Method,
			"path", path,
			"query", raw,
			"status", c.Writer.Status(),
			"latency_ms", latency.Milliseconds(),
			"client_ip", c.ClientIP(),
			"request_id", requestIDStr,
		)
	}
}

// Recovery recovers from panics and returns a proper error response.
func Recovery() gin.HandlerFunc {
	return gin.CustomRecovery(func(c *gin.Context, recovered interface{}) {
		requestID, _ := c.Get(RequestIDKey)
		requestIDStr := ""
		if id, ok := requestID.(string); ok {
			requestIDStr = id
		}

		logger.Error("Panic recovered",
			"error", recovered,
			"path", c.Request.URL.Path,
			"method", c.Request.Method,
			"request_id", requestIDStr,
		)

		c.JSON(500, gin.H{
			"error":      "Internal server error",
			"request_id": requestIDStr,
		})
		c.Abort()
	})
}

// CORS adds CORS headers to responses.
func CORS() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE, PATCH")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
