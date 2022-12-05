package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Greetings() gin.HandlerFunc {
	return func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"greetings": "hello world, api is good to go...",
		})
	}
}
