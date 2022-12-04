package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Greetings() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"greetings": "hello world, api is good to go...",
		})
	}
}
