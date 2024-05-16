package utils

import (
	"time"

	"github.com/sonochiwa/news/configs"

	"github.com/dgrijalva/jwt-go"
)

func GenerateJWT(username string) (string, error) {
	secret := configs.GlobalConfig.Auth.SecretKey

	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["username"] = username
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	tokenString, err := token.SignedString([]byte(secret))

	return tokenString, err
}
