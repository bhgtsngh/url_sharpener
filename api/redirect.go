package api

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"url_proj/storage"
)

func Redirect(c *gin.Context) {
	shortURL := c.Param("shortURL")
	originalURL, exists := storage.GetOriginalURL(shortURL)
	if !exists {
		c.JSON(http.StatusNotFound, gin.H{"error": "URL not found"})
		return
	}
	c.Redirect(http.StatusMovedPermanently, originalURL)
}
