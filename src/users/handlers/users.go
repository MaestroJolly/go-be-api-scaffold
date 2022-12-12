package handlers

import (
	"net/http"

	"github.com/MaestroJolly/go-be-api-scaffold/src/db/helpers"
	"github.com/gin-gonic/gin"
)

// function to handle authorized users
func AuthorizedUser() gin.HandlerFunc {
	return func(context *gin.Context) {
		user, err := helpers.CurrentUser(context)

		if err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		context.JSON(http.StatusOK, gin.H{
			"message": "authorized user",
			"data":    user,
		})
	}
}
