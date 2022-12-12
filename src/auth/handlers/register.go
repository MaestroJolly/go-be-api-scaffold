package handlers

import (
	"fmt"
	"net/http"

	"github.com/MaestroJolly/go-be-api-scaffold/src/db/models"
	"github.com/gin-gonic/gin"
)

type UserRegistration struct {
	UserName  string `json:"username" binding:"required"`
	FirstName string `json:"first_name" binding:"required"`
	LastName  string `json:"last_name" binding:"required"`
	Email     string `json:"email" binding:"required"`
	Password  string `json:"password" binding:"required"`
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

		user := models.User{
			Username:  data.UserName,
			FirstName: data.FirstName,
			LastName:  data.LastName,
			Email:     data.Email,
			Password:  data.Password,
		}

		savedUser, err := user.Save()

		if err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		context.JSON(http.StatusCreated, gin.H{
			"message": "registration successful",
			"user":    savedUser,
		})
	}
}
