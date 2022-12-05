package main

import (
	"github.com/gin-gonic/gin"

	authHandlers "github.com/MaestroJolly/go-be-api-scaffold/src/auth/handlers"
	greetingsHandlers "github.com/MaestroJolly/go-be-api-scaffold/src/greetings/handlers"
	healthHandlers "github.com/MaestroJolly/go-be-api-scaffold/src/health/handlers"
)

// Routes manager [Function to initiate routes]
func initRouter() *gin.Engine {
	r := gin.Default()

	// handler functions
	greetings := greetingsHandlers.Greetings()
	healthCheck := healthHandlers.HealthCheck()
	register := authHandlers.Register()

	// public routes group
	public := r.Group("/auth")

	r.GET("/", greetings)
	r.GET("/health", healthCheck)

	// public routes
	public.POST("/register", register)

	return r
}

// Main function
func main() {
	r := initRouter()
	r.Run(":8080")
}
