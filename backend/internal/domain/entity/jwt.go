package entity

import (
	"time"

	"github.com/caioap/ache-seu-pet/backend/config"
	"github.com/golang-jwt/jwt/v5"
)

type JWT struct {}

func Generate(userId int) (string, error) {
	jwtSecretKey := []byte(config.Config("JWT_SECRET"))
	token := jwt.New(jwt.SigningMethodEdDSA)
	claims := token.Claims.(jwt.MapClaims)
	claims["exp"] = time.Now().Add(24 * time.Hour)
	claims["authorized"] = true
	claims["user"] = userId
	tokenString, err := token.SignedString(jwtSecretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
