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

	// auth handler functions
	register := authHandlers.Register()
	login := authHandlers.Login()

	// public routes group
	public := r.Group("/auth")

	r.GET("/", greetings)
	r.GET("/health", healthCheck)

	// public routes
	public.POST("/register", register)
	public.POST("/login", login)

	return r
}

// Main function
func main() {
	r := initRouter()
	r.Run(":8080")
}
