package user

import (
	"github.com/accmi/words-api/app/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Credentials struct {
	Email string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// SignUp create users
func SignUpHandler(c *gin.Context) {
	var uc Credentials
	err := c.BindJSON(&uc)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})

		return
	}

	uc.Password, err = HashPassword(uc.Password)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})

		return
	}

	user := User{
		Email: uc.Email,
		PasswordsHash: uc.Password,
	}

	err = SaveUser(&user)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})

		return
	}

	err, user.Token = utils.CreateToken(user.Email)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, user)
}

// SignIn authenticate user
func SignInHandler(c *gin.Context) {
	var uc Credentials
	var uph string
	err := c.BindJSON(&uc)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})

		return
	}

	err = GetUserPasswordByEmail(uc.Email, &uph)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})

		return
	}

	res := CheckPasswordHash(uc.Password, uph)

	if res {
		err, token := utils.CreateToken(uc.Email)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})

			return
		}

		c.JSON(http.StatusOK, gin.H{
			"token": token,
		})

		return
	}

	c.JSON(http.StatusBadRequest, gin.H{
		"error": "Password isn't correct",
	})
}
