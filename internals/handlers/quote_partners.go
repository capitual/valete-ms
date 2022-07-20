package handlers

import (
	"net/http"

	"github.com/capitual/valete_ms/internals/dtos"
	"github.com/capitual/valete_ms/internals/repositories"
	"github.com/capitual/valete_ms/internals/services"
	"github.com/gin-gonic/gin"
)

func NewQuotePartner(c *gin.Context) {
	var dto dtos.QuotePartnerDto

	err := c.ShouldBindJSON(&dto)

	if err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	repository := &repositories.QuotePartnerRepository{}
	service := services.NewQuotePartnersService(repository)

	result, err := service.CreateQuotePartner(dto)

	if err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	c.JSON(http.StatusOK, successResponse(result))
}
