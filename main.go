package main

import (
	"github.com/gin-gonic/gin"
	"url_proj/api"
)

func main() {
	r := gin.Default()

	r.POST("/shorten", api.ShortenURL)
	r.GET("/:shortURL", api.Redirect)
	r.GET("/metrics", api.Metrics)

	r.Run(":8080") 
}
