package tests

import (
	"testing"
	"url_proj/api"
	"url_proj/storage"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestShortenURL(t *testing.T) {
	
	gin.SetMode(gin.TestMode)
	router := gin.Default()

	
	router.POST("/shorten", api.ShortenURL)

	
	tests := []struct {
		url          string
		expectedCode int
		expectedJSON string
	}{
		{
			url:          "https://youtube.com",
			expectedCode: 200,
			expectedJSON: `{"short_url":"7p0lP5"}`,
		},
		{
			url:          "https://udemy.com",
			expectedCode: 200,
			expectedJSON: `{"short_url":"fOMSvi"}`,
		},
		{
			url:          "https://wikipedia.com",
			expectedCode: 200,
			expectedJSON: `{"short_url":"zrRofb"}`, 
		}
		{
			url:          "https://google.com",
			expectedCode: 200,
			expectedJSON: `{"short_url":"mCywHS"}`, 
	}

	
	for _, tt := range tests {
		t.Run(tt.url, func(t *testing.T) {
			
			req := gin.CreateTestContext(nil)
			req.SetRequestBody([]byte(`{"url":"` + tt.url + `"}`))

			
			resp := req.DoRequest(router)

			
			assert.Equal(t, tt.expectedCode, resp.StatusCode)
			assert.JSONEq(t, tt.expectedJSON, resp.Body)
		})
	}
}
