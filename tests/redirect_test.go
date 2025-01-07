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

func TestRedirect(t *testing.T) {
	// Set up Gin for testing
	gin.SetMode(gin.TestMode)
	router := gin.Default()

	// Register the redirect route
	router.GET("/:shortURL", api.Redirect)

	// Mock data
	originalURL := "https://example.com"
	shortURL := "abc123"
	storage.SaveURL(originalURL, shortURL)

	// Test cases
	tests := []struct {
		shortURL        string
		expectedCode    int
		expectedLocation string
	}{
		{
			shortURL:        shortURL,
			expectedCode:    http.StatusMovedPermanently,
			expectedLocation: originalURL,
		},
		{
			shortURL:        "nonexistent",
			expectedCode:    http.StatusNotFound,
			expectedLocation: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.shortURL, func(t *testing.T) {
			// Create a GET request
			req := httptest.NewRequest(http.MethodGet, "/"+tt.shortURL, nil)
			resp := httptest.NewRecorder()

			// Perform the API call
			router.ServeHTTP(resp, req)

			// Validate response code
			assert.Equal(t, tt.expectedCode, resp.Code)

			// Validate the Location header for redirection
			if tt.expectedLocation != "" {
				assert.Equal(t, tt.expectedLocation, resp.Header().Get("Location"))
			}
		})
	}
}
