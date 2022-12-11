package helpers

import (
	"os"
	"strconv"
	"time"

	"github.com/MaestroJolly/go-be-api-scaffold/src/db/models"
	"github.com/golang-jwt/jwt/v4"
)

var privateKey = []byte(os.Getenv("JWT_PRIVATE_KEY"))

// function to generate the jwt token
func GenerateJWT(user models.User) (string, error) {
	tokenTTL, _ := strconv.Atoi(os.Getenv("TOKEN_TTL"))
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":    user.ID,
		"email": user.Email,
		"iat":   time.Now().Unix(),
		"eat":   time.Now().Add(time.Second * time.Duration(tokenTTL)).Unix(),
	})
	return token.SignedString(privateKey)
}
