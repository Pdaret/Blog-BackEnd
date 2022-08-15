package utils

import (
	"time"

	"github.com/dgrijalva/jwt-go/v4"
)

const SecretKey = "secret"

func GenerateJwt(issuer string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["iss"] = issuer
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()
	tokenString, err := token.SignedString([]byte(SecretKey))
	return tokenString, err
}

func ParseJwt(cookie string) (string, error) {
	token, err := jwt.Parse(cookie, func(token *jwt.Token) (interface{}, error) {
		return []byte(SecretKey), nil
	})
	if err != nil {
		return "", err
	}
	return token.Claims.(jwt.MapClaims)["iss"].(string), nil
}
