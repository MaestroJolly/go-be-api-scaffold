package handlers

import (
	"errors"
	"fmt"
	"net/http"

	helpers "github.com/MaestroJolly/go-be-api-scaffold/src/db/helpers"
	"github.com/MaestroJolly/go-be-api-scaffold/src/db/models"
	"github.com/gin-gonic/gin"
)

type UserLogin struct {
	ID       uint   `json:"id"`
	Email    string `json:"email"`
	Username string `json:"username"`
	Password string `json:"password" binding:"required"`
}

func Login() gin.HandlerFunc {
	return func(context *gin.Context) {
		data := UserLogin{}

		if err := context.ShouldBindJSON(&data); err != nil {
			fmt.Println(err)
			context.JSON(http.StatusBadRequest, gin.H{
				"error": fmt.Sprintf("%v", err),
			})
			return
		}

		if data.Email == "" && data.Password == "" {
			context.JSON(http.StatusBadRequest, gin.H{
				"error": fmt.Sprintf("%v", errors.New("email or username cannot be empty")),
			})
			return
		}

		user, err := models.FindUserByUsername(data.Username)

		if err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		err = user.ValidatePassword(data.Username)

		if err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		userJWT, err := helpers.GenerateJWT(user)

		if err != nil {
			fmt.Println(err)
			context.JSON(http.StatusBadRequest, gin.H{
				"error": fmt.Sprintf("%v", err),
			})
			return
		}

		context.JSON(http.StatusOK, gin.H{
			"message": "login successful",
			"data":    userJWT,
		})
	}
}
