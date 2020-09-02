package Controllers

import (
	"log"
	"net/http"

	Models "github.com/accmi/words-api/models"
	"github.com/gin-gonic/gin"
)

// GetUsers get all users
func GetUsers(c *gin.Context) {
	var users Models.Users

	err := users.GetAllUsers()

	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
	} else {
		c.JSON(http.StatusOK, users)
	}
}

// CreateUser add new user
func CreateUser(c *gin.Context) {
	user := Models.User{
		Name:  c.PostForm("name"),
		Email: c.PostForm("email"),
	}

	err := user.CreateUser()

	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		log.Panic(err)
	} else {
		c.JSON(http.StatusOK, user)
	}
}

// DeleteUser add new user
func DeleteUser(c *gin.Context) {
	var user Models.User

	id := c.Params.ByName("id")

	err := user.DeleteUser(id)

	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		log.Panic(err)
	} else {
		c.JSON(http.StatusOK, map[string]string{
			"isDeleted": "true",
		})
	}
}
