package routes

import (
	"github.com/capitual/valete_ms/integrations/controllers"
	internals_routes "github.com/capitual/valete_ms/internals/handlers"
	"github.com/capitual/valete_ms/routes/middlewares"
	"github.com/gin-gonic/gin"
)

func ConfigRoutes(router *gin.Engine) *gin.Engine {
	main := router.Group("api/v1")
	{
		pairs := main.Group("pairs")
		{
			pairs.GET("/", middlewares.Authorization(), controllers.GetPairs)
		}
		internals := main.Group("internals")
		{
			internals.POST("/categories", internals_routes.NewQuotaCategory)
			partners := internals.Group("partners")
			{
				partners.POST("/", internals_routes.NewPartner)
				partners.GET("/:id", middlewares.Secret(), internals_routes.GetPartnerById)
				partners.PUT("/:id", middlewares.Secret(), internals_routes.RevogatePartner)
				partners.GET("/", internals_routes.GetAllPartners)
			}
			quotasSettings := internals.Group("settings")
			{
				quotasSettings.POST("/", middlewares.Authorization(), internals_routes.NewQuotaSetting)
			}
		}
	}
	return router
}
