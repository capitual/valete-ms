package handlers

import (
	"net/http"

	"github.com/capitual/valete_ms/internals/dtos"
	"github.com/capitual/valete_ms/internals/repositories"
	"github.com/capitual/valete_ms/internals/services"
	"github.com/gin-gonic/gin"
)

func NewCurrency(c *gin.Context) {
	var dto dtos.CurrencyDto

	err := c.ShouldBindJSON(&dto)

	if err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	repository := &repositories.CurrencyRepository{}

	service := services.NewCurrencyService(repository)

	result, err := service.CreateCurrency(dto)

	if err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	c.JSON(http.StatusOK, successResponse(result))
}

func ListCurrencies(c *gin.Context) {
	search := c.Query("search")

	repository := &repositories.CurrencyRepository{}

	service := services.NewCurrencyService(repository)

	result, err := service.ListCurrencies(search)

	if err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	c.JSON(http.StatusOK, successResponse(result))
}
