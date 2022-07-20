package handlers

import (
	"net/http"

	"github.com/capitual/valete_ms/internals/dtos"
	"github.com/capitual/valete_ms/internals/repositories"
	"github.com/capitual/valete_ms/internals/services"
	"github.com/gin-gonic/gin"
)

func NewQuoteCategory(c *gin.Context) {
	var q dtos.QuoteCategoryDto

	err := c.ShouldBindJSON(&q)

	if err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	repository := &repositories.QuoteCategoriesRepository{}
	service := services.NewQuoteCategoryService(repository)

	result, err := service.CreateQuoteCategory(q)

	if err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	c.JSON(http.StatusOK, successResponse(result))
}
