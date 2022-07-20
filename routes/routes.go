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
			internals.POST("/categories", internals_routes.NewQuoteCategory)
			partners := internals.Group("partners")
			{
				partners.POST("/", internals_routes.NewPartner)
				partners.GET("/:id", middlewares.Secret(), internals_routes.GetPartnerById)
				partners.PUT("/revogate/:id", middlewares.Secret(), internals_routes.RevogatePartner)
				partners.GET("/", internals_routes.GetAllPartners)
			}
			quotesSettings := internals.Group("settings")
			{
				quotesSettings.POST("/", middlewares.Authorization(), internals_routes.NewQuoteSetting)
			}
			currencies := internals.Group("currencies")
			{
				currencies.POST("/", internals_routes.NewCurrency)
				currencies.GET("/", internals_routes.ListCurrencies)
			}
			quote_partners := internals.Group("quotepartners")
			{
				quote_partners.POST("/", internals_routes.NewQuotePartner)
			}
		}
	}

	// router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	// main.Run(":8080")
	return router
}
