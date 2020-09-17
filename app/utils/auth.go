package utils

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	"strings"
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

func TokenAuthMiddleware(c *gin.Context) {
	authorization := c.GetHeader("Authorization")
	st := strings.Fields(authorization)

	if len(st) != 2 {
		c.AbortWithStatus(http.StatusForbidden)
		log.Println("auth token isn't correct")
		return
	}

	if st[0] != "Bearer" {
		c.AbortWithStatus(http.StatusForbidden)
		log.Println("auth token isn't correct")
		return
	}

	res := Contains(CurrentTokens, st[1])

	if res < 0 {
		c.AbortWithStatus(http.StatusUnauthorized)
		log.Println("auth token isn't correct")
		return
	}

	c.Next()
}

func Contains(c []string, v string) int {
	for _, s := range c {
		if s == v {
			return 1
		}
	}

	return -1
}
