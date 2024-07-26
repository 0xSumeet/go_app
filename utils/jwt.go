package utils

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

var JWTSecret = []byte("your-secret-key")

func GenerateJWT(username, role string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": username,
		"role":     role,
		"exp":      time.Now().Add(time.Hour * 24).Unix(),
	})
	return token.SignedString(JWTSecret)
}
