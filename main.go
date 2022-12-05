package main

import (
	"github.com/gin-gonic/gin"

	greetingsHandlers "github.com/MaestroJolly/go-be-api-scaffold/src/greetings/handlers"
	healthHandlers "github.com/MaestroJolly/go-be-api-scaffold/src/health/handlers"
)

// Routes manager [Function to initiate routes]
func initRouter() *gin.Engine {
	r := gin.Default()

	// handler functions
	greetings := greetingsHandlers.Greetings()
	healthCheck := healthHandlers.HealthCheck()

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
