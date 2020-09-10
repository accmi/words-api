package user

import (
	"fmt"
	"github.com/accmi/words-api/app/utils"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type Credentials struct {
	Email string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// SignUp create users
func SignUpHandler(c *gin.Context) {
	password := c.PostForm("password")
	hash, err := HashPassword(password)

	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		log.Panic(err)
		return
	}

	user := User{
		PasswordsHash: hash,
	}

	err = c.BindJSON(&user)

	err = SaveUser(&user)

	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		log.Panic(err)
		return
	}

	err = utils.CreateToken(user.Token)

	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		log.Panic(err)
		return
	}

	user.Token = utils.CurrentToken

	c.JSON(http.StatusOK, user)
}

// SignIn authenticate user
func SignInHandler(c *gin.Context) {
	var uc Credentials
	err := c.BindJSON(&uc)

	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		log.Panic(err)
		return
	}

	fmt.Println(uc)
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
