package handlers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserRegistration struct {
	UserName string `json:"user_name" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func Register() gin.HandlerFunc {
	return func(context *gin.Context) {
		data := UserRegistration{}

		if err := context.ShouldBindJSON(&data); err != nil {
			fmt.Println(err)
			context.JSON(http.StatusBadRequest, gin.H{
				"error": fmt.Sprintf("%v", err),
			})
			return
		}

		context.JSON(http.StatusOK, gin.H{
			"message": "registration successful",
			"data":    data,
		})
	}
}
