package main

import (
	"github.com/gin-gonic/gin"

	"github.com/MaestroJolly/go-be-api-scaffold/src/greetings/handlers"
)

// Routes manager [Function to initiate routes]
func initRouter() *gin.Engine {
	r := gin.Default()

	// handler functions
	greetings := handlers.Greetings()
	healthCheck := handlers.HealthCheck()

	// routes
	r.GET("/", greetings)
	r.GET("/api/health", healthCheck)

	return r
}

// Main function
func main() {
	r := initRouter()
	r.Run(":8080")
}
