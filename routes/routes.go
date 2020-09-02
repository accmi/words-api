package routes

import (
	Controllers "github.com/accmi/words-api/controllers"
	"github.com/gin-gonic/gin"
)

// SetupRouter contains all routers
func SetupRouter() *gin.Engine {
	r := gin.Default()
	grp := r.Group("users")
	{
		grp.GET("", Controllers.GetUsers)
		grp.POST("", Controllers.CreateUser)
		grp.DELETE(":id", Controllers.DeleteUser)
	}

	return r
}
