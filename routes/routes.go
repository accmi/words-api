package routes

import (
	Controllers "github.com/accmi/words-api/controllers"
	"github.com/gin-gonic/gin"
)

// SetupRouter contains all routers
func SetupRouter() *gin.Engine {
	r := gin.Default()
	grp := r.Group("authentication")
	{
		//grp.GET("", Controllers.GetUsers)
		//grp.POST("", Controllers.CreateUser)
		//grp.DELETE(":id", Controllers.DeleteUser)

		grp.POST("up", Controllers.SignUp)
		grp.POST("in", Controllers.SignIn)
	}

	return r
}
