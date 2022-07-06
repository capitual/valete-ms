package controllers

import (
	"github.com/capitual/valete_ms/helpers"
	"github.com/gin-gonic/gin"
)

func GetPairs(c *gin.Context) {
	result, err := helpers.OwsRequest("GET", "pairs", "")

	if err != nil {
		c.JSON(400, gin.H{
			"error": err,
		})
		return
	}

	c.JSON(200, result)
}
