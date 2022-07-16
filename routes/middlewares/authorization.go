package middlewares

import (
	"net/http"

	infra "github.com/capitual/valete_ms/infra"
	"github.com/capitual/valete_ms/internals/models"
	"github.com/gin-gonic/gin"
)

func Authorization() gin.HandlerFunc {
	return func(c *gin.Context) {
		partner_key := c.GetHeader("x-client-key")
		partner_id := c.GetHeader("x-client-id")

		if partner_key == "" || partner_id == "" {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		db := infra.GetDatabase()

		var p models.Partner

		err := db.First(&p, `partner_key = ? AND partner_id = ?
		"AND active = true`, partner_key, partner_id).Error

		if err != nil {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		c.Set("partner", p.ID)

		c.Next()
	}
}
