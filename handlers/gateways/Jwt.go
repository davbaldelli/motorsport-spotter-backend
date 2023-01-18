package gateways

import (
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"time"
)

func GenerateJWT(username, role string, secret string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"authorized": true,
		"email":      username,
		"role":       role,
		"exp":        time.Now().Add(time.Hour * 24).Unix(),
	})

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString([]byte(secret))

	if err != nil {
		return "", fmt.Errorf("Something Went Wrong: %s ", err.Error())
	}
	return tokenString, nil
}
