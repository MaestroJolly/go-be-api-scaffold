package main

import (
	"github.com/gin-gonic/gin"

	authHandlers "github.com/MaestroJolly/go-be-api-scaffold/src/auth/handlers"
	greetingsHandlers "github.com/MaestroJolly/go-be-api-scaffold/src/greetings/handlers"
	healthHandlers "github.com/MaestroJolly/go-be-api-scaffold/src/health/handlers"
	"github.com/MaestroJolly/go-be-api-scaffold/src/middlewares"
	userHandlers "github.com/MaestroJolly/go-be-api-scaffold/src/users/handlers"
)

// Routes manager [Function to initiate routes]
func initRouter() *gin.Engine {
	router := gin.Default()

	// handler functions
	greetings := greetingsHandlers.Greetings()
	healthCheck := healthHandlers.HealthCheck()

	// auth handler functions
	register := authHandlers.Register()
	login := authHandlers.Login()

	// user handler functions
	user := userHandlers.AuthorizedUser()

	// public router group
	public := router.Group("/auth")

	router.GET("/", greetings)
	router.GET("/health", healthCheck)

	// public router
	public.POST("/register", register)
	public.POST("/login", login)

	// protected router group
	protected := router.Group("/api")
	protected.Use(middlewares.IsAuthorized())

	protected.POST("/users", user)

	return router
}

// Main function
func main() {
	r := initRouter()
	r.Run(":8080")
}
