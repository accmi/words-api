package controllers

import (
	"github.com/accmi/words-api/models"
	"github.com/accmi/words-api/repositories"
	"github.com/gin-gonic/gin"
	// jwt "github.com/appleboy/gin-jwt/v2"
	"log"
	"net/http"
)

// SignUp create users
func SignUp(c *gin.Context) {
	password := c.PostForm("password")
	hash, err := models.HashPassword(password)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		log.Panic(err)
		return
	}
	// claims := jwt.ExtractClaims(c)
	user := models.User{
		Email: c.PostForm("email"),
		PasswordsHash: hash,
		// Token:
	}

	err = repositories.CreateUser(&user)
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		log.Panic(err)
		return
	}

	c.JSON(http.StatusOK, user)
}

// SignIn authenticate user
func SignIn(c *gin.Context) {
	//password := c.PostForm("password")
	//email := c.PostForm("email")
}

//// GetUsers get all users
//func GetUsers(c *gin.Context) {
//	var users []models.User
//
//	err := models.GetUsers(&users)
//
//	if err != nil {
//		c.AbortWithStatus(http.StatusBadRequest)
//	} else {
//		c.JSON(http.StatusOK, users)
//	}
//}
//
//// CreateUser add new user
//func CreateUser(c *gin.Context) {
//	user := models.User{
//		Name:  c.PostForm("name"),
//		Email: c.PostForm("email"),
//	}
//
//	err := user.CreateUser()
//
//	if err != nil {
//		c.AbortWithStatus(http.StatusBadRequest)
//		log.Panic(err)
//	} else {
//		c.JSON(http.StatusOK, user)
//	}
//}
//
//// DeleteUser add new user
//func DeleteUser(c *gin.Context) {
//	var user models.User
//
//	ids := c.Params.ByName("id")
//	id, err := strconv.Atoi(ids)
//	if err != nil {
//		c.AbortWithStatus(http.StatusBadRequest)
//		log.Panic(err)
//	}
//
//	err = user.DeleteUser(id)
//
//	if err != nil {
//		if err.Error() == "not found" {
//			c.JSON(http.StatusOK, map[string]string{
//				"error": "this user doesn't exist",
//			})
//			return
//		}
//		c.AbortWithStatus(http.StatusNotFound)
//		log.Panic(err)
//	} else {
//		c.JSON(http.StatusOK, map[string]string{
//			"isDeleted": "true",
//		})
//	}
//}
