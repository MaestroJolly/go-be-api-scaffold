package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// function to handle authorized users
func AuthorizedUser() gin.HandlerFunc {
	return func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"message": "authorized user",
		})
	}
}
