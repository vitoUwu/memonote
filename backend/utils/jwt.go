package utils

import (
	"log"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type UserData struct {
	UserID   string `json:"user_id"`
	Username string `json:"username"`
	jwt.RegisteredClaims
}

func EncodeUserToken(id string, username string) (string, error) {
	data := UserData{
		UserID:   id,
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, data)

	secret := os.Getenv("SECRET")
	if secret == "" {
		log.Fatal("SECRET is missing")
	}

	return token.SignedString([]byte(secret))
}

func DecodeUserToken(token string) (*UserData, error) {
	secret := os.Getenv("SECRET")
	if secret == "" {
		log.Fatal("SECRET is missing")
	}

	data, err := jwt.ParseWithClaims(
		token,
		&UserData{},
		func(t *jwt.Token) (interface{}, error) {
			return []byte(secret), nil
		},
	)

	if err != nil {
		return nil, err
	}

	return data.Claims.(*UserData), nil
}
