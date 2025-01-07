package tests

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"url_proj/api"
	"url_proj/storage"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestMetrics(t *testing.T) {
	// Set up Gin for testing
	gin.SetMode(gin.TestMode)
	router := gin.Default()

	// Register the metrics route
	router.GET("/metrics", api.Metrics)

	// Simulate shortening URLs with various domains
	storage.SaveURL("https://youtube.com", "abc123")
	storage.SaveURL("https://udemy.com", "xyz456")
	storage.SaveURL("https://wikipedia.com", "def789")
	storage.SaveURL("https://udemy.com", "ghi012")
	storage.SaveURL("https://google.com", "jkl345")
	storage.SaveURL("https://udemy.com", "mno678")
	storage.SaveURL("https://youtube.com", "pqr901")

	// Expected response
	expectedJSON := `{"top_domains":{"udemy.com":3,"youtube.com":2,"wikipedia.com":1}}`

	// Test the metrics API
	req := httptest.NewRequest(http.MethodGet, "/metrics", nil)
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	// Validate response
	assert.Equal(t, http.StatusOK, resp.Code)
	assert.JSONEq(t, expectedJSON, resp.Body.String())
}
