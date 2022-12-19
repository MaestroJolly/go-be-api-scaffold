package middlewares

import (
	"io"
	"os"

	"github.com/gin-gonic/gin"
)

func RouteInitializer() *gin.Engine {

	router := gin.Default()

	ginMode := os.Getenv("GIN_MODE")
	if ginMode == "local" {
		router = gin.New()

		// Logging to a file.
		file, _ := os.Create("debug.log")
		gin.DefaultWriter = io.MultiWriter(file)
	}

	return router
}
