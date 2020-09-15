package user

import (
	"github.com/gin-gonic/gin"
)

func CreateRoutes(ge *gin.Engine) {
	grp := ge.Group("auth")
	{
		grp.POST("up", SignUpHandler)
		grp.POST("in", SignInHandler)
	}
}
