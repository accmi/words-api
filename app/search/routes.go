package search

import (
	"github.com/gin-gonic/gin"
)

func CreateRoutes(ge *gin.Engine) {
	// ge.Use(utils.TokenAuthMiddleware)

	grp := ge.Group("search")
	{
		grp.GET("", FindParticularEntity)
	}
}
