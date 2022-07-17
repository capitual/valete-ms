package handlers

import (
	"net/http"

	"github.com/capitual/valete_ms/internals/dtos"
	"github.com/capitual/valete_ms/internals/repositories"
	"github.com/capitual/valete_ms/internals/services"
	"github.com/gin-gonic/gin"
)

func NewQuotaCategory(c *gin.Context) {
	var q dtos.QuotaCategoryDto

	err := c.ShouldBindJSON(&q)

	if err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	repository := &repositories.QuotaCategoriesRepository{}
	service := services.NewQuotaCategoryService(repository)

	result, err := service.CreateQuotaCategory(q)

	if err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	c.JSON(http.StatusOK, successResponse(result))
}
