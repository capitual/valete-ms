package routes

import (
	"github.com/capitual/valete_ms/integrations/controllers"
	"github.com/gin-gonic/gin"
)

func ConfigRoutes(router *gin.Engine) *gin.Engine {
	main := router.Group("api/v1")
	{
		pairs := main.Group("pairs")
		{
			pairs.GET("/", controllers.GetPairs)
		}
	}
	return router
}
