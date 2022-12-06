package handlers

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/MaestroJolly/go-be-api-scaffold/src/db/models"
	jwtGenerator "github.com/MaestroJolly/go-be-api-scaffold/src/helpers"
	"github.com/gin-gonic/gin"
)

type UserLogin struct {
	ID       string `json:"id"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

var dummyLoginData UserLogin = UserLogin{
	ID:       "156337",
	Email:    "example@example.com",
	Password: "password",
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

		if data.Email != dummyLoginData.Email || data.Password != dummyLoginData.Password {
			context.JSON(http.StatusBadRequest, gin.H{
				"error": fmt.Sprintf("%v", errors.New("email or password is incorrect")),
			})
			return
		}

		user := models.User{}

		user.ID = dummyLoginData.ID
		user.Email = dummyLoginData.Email

		userJWT, err := jwtGenerator.GenerateJWT(user)

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
