package handlers

import (
	"net/http"

	"github.com/capitual/valete_ms/internals/dtos"
	"github.com/capitual/valete_ms/internals/repositories"
	"github.com/capitual/valete_ms/internals/services"
	"github.com/gin-gonic/gin"
)

func NewQuotaSetting(c *gin.Context) {
	var dto dtos.QuotaSettingDto

	err := c.ShouldBindJSON(&dto)

	if err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	partner, _ := c.Get("partner")
	partner_id, _ := partner.(uint)

	dto.PartnerId = partner_id

	repository := &repositories.QuotaSettingRepository{}
	service := services.NewQuotaSettingService(repository)

	result, err := service.CreateQuotaSetting(dto)

	if err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	c.JSON(http.StatusOK, successResponse(result))

}
