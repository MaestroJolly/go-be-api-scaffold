package helpers

import (
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

var privateKey = []byte(os.Getenv("JWT_PRIVATE_KEY"))

// function to validate the jwt token
func ValidateJWT(context *gin.Context) error {
	token, err := GetToken(context)
	if err != nil {
		return err
	}
	_, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		return nil
	}
	return errors.New("invalid token provided")

}

// function to get the token from the context
func GetToken(context *gin.Context) (*jwt.Token, error) {
	tokenString, err := getTokenFromRequest(context)

	if err != nil {
		return nil, err
	}

	token, err := jwt.Parse(*tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return privateKey, nil
	})
	return token, err
}

// function to get the token from request
func getTokenFromRequest(context *gin.Context) (*string, error) {
	bearerToken := context.Request.Header.Get("Authorization")
	splitToken := strings.Split(bearerToken, " ")
	if len(splitToken) == 2 {
		return &splitToken[1], nil
	}
	return nil, errors.New("Authentication required")
}
