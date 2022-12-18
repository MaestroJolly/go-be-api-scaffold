package middlewares

import (
	"net/http"

	"github.com/MaestroJolly/go-be-api-scaffold/src/db/models"
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

		userId := helpers.GetUserIdFromToken(context)
		loginhash, err := models.FindLoginHashByUserId(userId)

		if err != nil {
			context.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			context.Abort()
			return
		}

		if loginhash.Hash == nil {
			context.JSON(http.StatusUnauthorized, gin.H{"error": "User already logged out. Please log in again."})
			context.Abort()
			return
		}

		context.Next()
	}
}
