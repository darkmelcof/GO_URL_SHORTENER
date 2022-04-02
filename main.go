package main

import (
	"fmt"

	"github.com/darkmelcof/go-url-shortener/handler"
	"github.com/darkmelcof/go-url-shortener/store"
	"github.com/gin-gonic/gin" // HTTP Web Framework
)

const port string = "9808"

func main() {

	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hey Go URL Shortener !",
		})
	})

	// Endpoint
	r.POST("/create-short-url", func(c *gin.Context) {
		handler.CreateShortUrl(c)
	})

	// Endpoint
	r.GET("/:shortUrl", func(c *gin.Context) {
		handler.HandleShortUrlRedirect(c)
	})

	// Note that store initialization happens here
	store.InitializeStore()

	err := r.Run(":" + port)
	if err != nil {
		panic(fmt.Sprintf("Failed to start the web server - Error: %v", err))
	}

}
