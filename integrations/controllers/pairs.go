package controllers

import (
	"github.com/capitual/valete_ms/integrations"
	"github.com/gin-gonic/gin"
)

func GetPairs(c *gin.Context) {
	result, err := integrations.OwsRequest("GET", "pairs", "")

	if err != nil {
		c.JSON(400, gin.H{
			"error": err,
		})
		return
	}

	c.JSON(200, result)
}
