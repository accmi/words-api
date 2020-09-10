package user

import (
	"github.com/gin-gonic/gin"
)

func CreateRoutes(ge *gin.Engine) {
	grp := ge.Group("authentication")
	{
		//grp.GET("", Controllers.GetUsers)
		//grp.POST("", Controllers.CreateUser)
		//grp.DELETE(":id", Controllers.DeleteUser)

		grp.POST("up", SignUpHandler)
		grp.POST("in", SignInHandler)
	}
}