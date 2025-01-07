package api

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"url_proj/storage"
	"math/rand"
	"time"
)

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

func generateShortURL() string {
	rand.Seed(time.Now().UnixNano())
	b := make([]rune, 6)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func ShortenURL(c *gin.Context) {
	type Request struct {
		URL string `json:"url"`
	}

	var req Request
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	if shortURL, exists := storage.GetShortURL(req.URL); exists {
		c.JSON(http.StatusOK, gin.H{"short_url": shortURL})
		return
	}

	shortURL := generateShortURL()
	storage.SaveURL(req.URL, shortURL)
	c.JSON(http.StatusOK, gin.H{"short_url": shortURL})
}
