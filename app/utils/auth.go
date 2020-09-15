package utils

import (
	"os"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

var CurrentTokens []string

func CreateToken(uId string) (error, string) {
	var err error = nil

	atClaims := jwt.MapClaims{
		"authorized": true,
		"user_id": uId,
		"exp": time.Now().Add(time.Hour * 1).Unix(),
	}

	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	secret := os.Getenv("SECRET_KEY")
	token, err := at.SignedString([]byte(os.Getenv(secret)))

	CurrentTokens = append(CurrentTokens, token)

	return err, token
}
