package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Routes manager [Function to initiate routes]
func initRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "api is up and running...",
		})
	})

	return r
}

// Main function
func main() {
	r := initRouter()
	r.Run(":8080")
}
