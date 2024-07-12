package utils

import (
	"time"

	"github.com/golang-jwt/jwt"
)

func GenerateToken(email string, userId int64) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": email,
		"exp":   time.Now().Add(time.Hour * 2).Unix(),
		"userId": userId,
	})

	tokenString, err := token.SignedString([]byte("secret"))
	if err != nil {
		// handle the error, e.g. log or return an empty string
		return ""
	}
	return tokenString
}