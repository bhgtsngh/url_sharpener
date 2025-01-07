package api

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"url_proj/storage"
)

func Metrics(c *gin.Context) {
	topDomains := storage.GetTopDomains(3)
	c.JSON(http.StatusOK, gin.H{"top_domains": topDomains,})
}
