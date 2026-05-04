// Package handler provides HTTP handlers for the API endpoints.
package handler

import (
	"net/http"

	"api/internal/model"
	"api/internal/service"
	"api/pkg/logger"

	"github.com/gin-gonic/gin"
)

// Handler holds dependencies for HTTP handlers.
type Handler struct {
	service service.DataService
}

// NewHandler creates a new handler instance.
func NewHandler(svc service.DataService) *Handler {
	return &Handler{
		service: svc,
	}
}

// GetAllData handles GET / requests to return all data.
func (h *Handler) GetAllData(c *gin.Context) {
	data := h.service.GetAllData()
	c.JSON(http.StatusOK, data)
}

// GetDataByID handles GET /:guid requests to return data by GUID.
func (h *Handler) GetDataByID(c *gin.Context) {
	guid := c.Param("guid")

	// Validate GUID format
	if !model.ValidateGUID(guid) {
		requestID, _ := c.Get("request_id")
		logger.Warn("Invalid GUID format", "guid", guid, "request_id", requestID)
		c.JSON(http.StatusBadRequest, gin.H{
			"error":      "Invalid GUID format",
			"request_id": requestID,
		})
		return
	}

	data := h.service.GetDataByGUID(guid)
	if data == nil {
		requestID, _ := c.Get("request_id")
		c.JSON(http.StatusNotFound, gin.H{
			"error":      "Data not found",
			"request_id": requestID,
		})
		return
	}

	c.JSON(http.StatusOK, data)
}

// HealthCheck handles GET /health requests for health checks.
func (h *Handler) HealthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status": "healthy",
	})
}
