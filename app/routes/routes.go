package routes

import (
	"github.com/accmi/words-api/app/user"
	"github.com/gin-gonic/gin"
)

// SetupRouter contains all routers
func SetupRouter() *gin.Engine {
	r := gin.Default()

	user.CreateRoutes(r)

	return r
}
