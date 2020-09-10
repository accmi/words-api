package utils

import (
	"os"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

var CurrentToken string

func CreateToken(uId string) error {
	var err error = nil

	atClaims := jwt.MapClaims{
		"authorized": true,
		"user_id": uId,
		"exp": time.Now().Add(time.Hour * 1).Unix(),
	}

	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	secret := os.Getenv("SECRET_KEY")
	CurrentToken, err = at.SignedString([]byte(os.Getenv(secret)))

	return err
}
