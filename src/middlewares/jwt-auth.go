package middlewares

import (
	"net/http"

	helpers "github.com/MaestroJolly/go-be-api-scaffold/src/helpers"
	"github.com/gin-gonic/gin"
)

// Jwt auth token middlewares validator
func IsAuthorized() gin.HandlerFunc {
	return func(context *gin.Context) {
		err := helpers.ValidateJWT(context)
		if err != nil {
			context.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			context.Abort()
			return
		}
		context.Next()
	}
}
