package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	authHandlers "github.com/MaestroJolly/go-be-api-scaffold/src/auth/handlers"
	greetingsHandlers "github.com/MaestroJolly/go-be-api-scaffold/src/greetings/handlers"
	healthHandlers "github.com/MaestroJolly/go-be-api-scaffold/src/health/handlers"
	"github.com/MaestroJolly/go-be-api-scaffold/src/middlewares"
	userHandlers "github.com/MaestroJolly/go-be-api-scaffold/src/users/handlers"

	database "github.com/MaestroJolly/go-be-api-scaffold/src/db"
)

// function to initiate environment variables
func initEnv() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

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
	initEnv()
	database.Connect()
	router := initRouter()
	router.Run(":8080")
}
