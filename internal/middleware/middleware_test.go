package middleware

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func setupTestRouter() *gin.Engine {
	gin.SetMode(gin.TestMode)
	return gin.New()
}

func TestRequestID(t *testing.T) {
	router := setupTestRouter()
	router.Use(RequestID())
	router.GET("/test", func(c *gin.Context) {
		requestID, exists := c.Get(RequestIDKey)
		assert.True(t, exists)
		assert.NotEmpty(t, requestID)
		c.JSON(200, gin.H{"request_id": requestID})
	})

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/test", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.NotEmpty(t, w.Header().Get(RequestIDHeader))
}

func TestRequestID_ExistingHeader(t *testing.T) {
	router := setupTestRouter()
	router.Use(RequestID())
	router.GET("/test", func(c *gin.Context) {
		requestID, _ := c.Get(RequestIDKey)
		c.JSON(200, gin.H{"request_id": requestID})
	})

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/test", nil)
	req.Header.Set(RequestIDHeader, "custom-request-id")
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "custom-request-id", w.Header().Get(RequestIDHeader))
}

func TestCORS(t *testing.T) {
	router := setupTestRouter()
	router.Use(CORS())
	router.GET("/test", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/test", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, "*", w.Header().Get("Access-Control-Allow-Origin"))
	assert.Equal(t, "true", w.Header().Get("Access-Control-Allow-Credentials"))
}

func TestCORS_OPTIONS(t *testing.T) {
	router := setupTestRouter()
	router.Use(CORS())
	router.OPTIONS("/test", func(_ *gin.Context) {})

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("OPTIONS", "/test", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 204, w.Code)
}
