package handler

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"api/internal/model"
	"api/internal/service"
)

var testGUID = "05024756-765e-41a9-89d7-1407436d9a58"

var testData = []model.Data{
	{
		GUID:       testGUID,
		School:     "Iowa State University",
		Mascot:     "Cy the Cardinal",
		Nickname:   "Cyclones",
		Location:   "Ames, IA, USA",
		LatLong:    "42.026111,-93.648333",
		NCAA:       "Division I",
		Conference: "Big 12 Conference",
	},
}

// setupTestRouter creates a test router with a mock service.
func setupTestRouter() (*gin.Engine, *Handler) {
	gin.SetMode(gin.TestMode)

	// Create a temporary service with test data
	svc := &mockService{
		data: testData,
	}

	h := NewHandler(svc)
	router := gin.New()
	router.GET("/health", h.HealthCheck)
	router.GET("/", h.GetAllData)
	router.GET("/:guid", h.GetDataByID)

	return router, h
}

// mockService is a mock implementation of the service for testing.
type mockService struct {
	data []model.Data
}

func (m *mockService) GetAllData() []model.Data {
	result := make([]model.Data, len(m.data))
	copy(result, m.data)
	return result
}

func (m *mockService) GetDataByGUID(guid string) *model.Data {
	for i := range m.data {
		if m.data[i].GUID == guid {
			result := m.data[i]
			return &result
		}
	}
	return nil
}

func TestHealthCheck(t *testing.T) {
	router, _ := setupTestRouter()

	tests := []struct {
		name           string
		method         string
		path           string
		expectedStatus int
		expectedBody   map[string]interface{}
	}{
		{
			name:           "health check returns 200",
			method:         "GET",
			path:           "/health",
			expectedStatus: http.StatusOK,
			expectedBody: map[string]interface{}{
				"status": "healthy",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			req, _ := http.NewRequest(tt.method, tt.path, nil)
			router.ServeHTTP(w, req)

			assert.Equal(t, tt.expectedStatus, w.Code)

			var result map[string]interface{}
			err := json.Unmarshal(w.Body.Bytes(), &result)
			require.NoError(t, err)
			assert.Equal(t, tt.expectedBody["status"], result["status"])
		})
	}
}

func TestGetAllData(t *testing.T) {
	router, _ := setupTestRouter()

	tests := []struct {
		name           string
		method         string
		path           string
		expectedStatus int
		validateFunc   func(t *testing.T, body []byte)
	}{
		{
			name:           "get all data returns 200",
			method:         "GET",
			path:           "/",
			expectedStatus: http.StatusOK,
			validateFunc: func(t *testing.T, body []byte) {
				var result []model.Data
				err := json.Unmarshal(body, &result)
				require.NoError(t, err)
				assert.Greater(t, len(result), 0)
				assert.Equal(t, testData[0].GUID, result[0].GUID)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			req, _ := http.NewRequest(tt.method, tt.path, nil)
			router.ServeHTTP(w, req)

			assert.Equal(t, tt.expectedStatus, w.Code)
			if tt.validateFunc != nil {
				tt.validateFunc(t, w.Body.Bytes())
			}
		})
	}
}

func TestGetDataByID(t *testing.T) {
	router, _ := setupTestRouter()

	tests := []struct {
		name           string
		method         string
		path           string
		expectedStatus int
		validateFunc   func(t *testing.T, body []byte)
	}{
		{
			name:           "get data by valid GUID returns 200",
			method:         "GET",
			path:           "/" + testGUID,
			expectedStatus: http.StatusOK,
			validateFunc: func(t *testing.T, body []byte) {
				var result model.Data
				err := json.Unmarshal(body, &result)
				require.NoError(t, err)
				assert.Equal(t, testData[0], result)
			},
		},
		{
			name:           "get data by invalid GUID format returns 400",
			method:         "GET",
			path:           "/this-is-a-bad-guid",
			expectedStatus: http.StatusBadRequest,
			validateFunc: func(t *testing.T, body []byte) {
				var result map[string]interface{}
				err := json.Unmarshal(body, &result)
				require.NoError(t, err)
				assert.Equal(t, "Invalid GUID format", result["error"])
			},
		},
		{
			name:           "get data by non-existent GUID returns 404",
			method:         "GET",
			path:           "/00000000-0000-0000-0000-000000000000",
			expectedStatus: http.StatusNotFound,
			validateFunc: func(t *testing.T, body []byte) {
				var result map[string]interface{}
				err := json.Unmarshal(body, &result)
				require.NoError(t, err)
				assert.Equal(t, "Data not found", result["error"])
			},
		},
		{
			name:           "get data by invalid GUID length returns 400",
			method:         "GET",
			path:           "/short",
			expectedStatus: http.StatusBadRequest,
			validateFunc: func(t *testing.T, body []byte) {
				var result map[string]interface{}
				err := json.Unmarshal(body, &result)
				require.NoError(t, err)
				assert.Equal(t, "Invalid GUID format", result["error"])
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			req, _ := http.NewRequest(tt.method, tt.path, nil)
			router.ServeHTTP(w, req)

			assert.Equal(t, tt.expectedStatus, w.Code)
			if tt.validateFunc != nil {
				tt.validateFunc(t, w.Body.Bytes())
			}
		})
	}
}

func TestValidateGUID(t *testing.T) {
	tests := []struct {
		name     string
		guid     string
		expected bool
	}{
		{
			name:     "valid GUID",
			guid:     "05024756-765e-41a9-89d7-1407436d9a58",
			expected: true,
		},
		{
			name:     "invalid GUID - too short",
			guid:     "short",
			expected: false,
		},
		{
			name:     "invalid GUID - wrong format",
			guid:     "this-is-a-bad-guid",
			expected: false,
		},
		{
			name:     "invalid GUID - missing dashes",
			guid:     "05024756765e41a989d71407436d9a58",
			expected: false,
		},
		{
			name:     "invalid GUID - wrong dash positions",
			guid:     "0502475-6765e-41a9-89d7-1407436d9a58",
			expected: false,
		},
		{
			name:     "valid GUID uppercase",
			guid:     "05024756-765E-41A9-89D7-1407436D9A58",
			expected: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := model.ValidateGUID(tt.guid)
			assert.Equal(t, tt.expected, result)
		})
	}
}

// Integration test with real service
func TestHandlerIntegration(t *testing.T) {
	// This test requires data.json to exist
	svc, err := service.NewService("./data.json")
	if err != nil {
		t.Skipf("Skipping integration test: %v", err)
		return
	}

	h := NewHandler(svc)
	router := gin.New()
	router.GET("/", h.GetAllData)
	router.GET("/:guid", h.GetDataByID)

	t.Run("integration get all data", func(t *testing.T) {
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/", nil)
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)
		var result []model.Data
		err := json.Unmarshal(w.Body.Bytes(), &result)
		require.NoError(t, err)
		assert.Greater(t, len(result), 0)
	})

	t.Run("integration get data by GUID", func(t *testing.T) {
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/"+testGUID, nil)
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)
		var result model.Data
		err := json.Unmarshal(w.Body.Bytes(), &result)
		require.NoError(t, err)
		assert.Equal(t, testGUID, result.GUID)
	})
}
