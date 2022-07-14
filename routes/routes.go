package routes

import (
	"github.com/capitual/valete_ms/integrations/controllers"
	internals_routes "github.com/capitual/valete_ms/internals/handlers"
	"github.com/gin-gonic/gin"
)

func ConfigRoutes(router *gin.Engine) *gin.Engine {
	main := router.Group("api/v1")
	{
		pairs := main.Group("pairs")
		{
			pairs.GET("/", controllers.GetPairs)
		}
		internals := main.Group("internals")
		{
			internals.POST("/categories", internals_routes.NewQuoteCategory)
			internals.POST("/partners", internals_routes.NewPartner)

		}
	}
	return router
}
