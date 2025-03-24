package utils

import (
	"oauth/models"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func GenerateJWT(user models.User) (string, error) {
	var mySigningKey = []byte("secret")

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": user.Username,
		"exp":      time.Now().Add(time.Hour * 24).Unix(),
	})

	tokenString, err := token.SignedString(mySigningKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
