package handlers

import (
	"errors"
	"fmt"
	"net/http"

	authhelpers "github.com/MaestroJolly/go-be-api-scaffold/src/auth/helpers"
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

		if data.Email == "" && data.Username == "" {
			context.JSON(http.StatusBadRequest, gin.H{
				"error": fmt.Sprintf("%v", errors.New("email or username cannot be empty")),
			})
			return
		}

		username := data.Username

		if username == "" {
			username = data.Email
		}

		user, err := models.FindUserByUsernameOrEmail(username)

		if err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if user.Email == "" {
			context.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
			return
		}

		err = user.ValidatePassword(data.Password)

		if err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"error": authhelpers.ErrorMessageNormalizer[err.Error()]})
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
			"message": "successfully logged in",
			"data":    userJWT,
		})
	}
}
