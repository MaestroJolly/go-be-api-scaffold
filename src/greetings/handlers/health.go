package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func HealthCheck() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "api is up and running...",
		})
	}
}
