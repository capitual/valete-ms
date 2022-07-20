package handlers

import (
	"net/http"

	"github.com/capitual/valete_ms/internals/dtos"
	"github.com/capitual/valete_ms/internals/repositories"
	"github.com/capitual/valete_ms/internals/services"
	"github.com/gin-gonic/gin"
)

func NewQuoteSetting(c *gin.Context) {
	var dto dtos.QuoteSettingDto

	err := c.ShouldBindJSON(&dto)

	if err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	partner, _ := c.Get("partner")
	partner_id, ok := partner.(uint)

	// switch partner.(type) {
	// 	case uint, uint32, uint64, int, int16, int32, int64 :
	// 	case string:
	// }

	if !ok {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	dto.PartnerId = partner_id

	repository := &repositories.QuoteSettingRepository{}
	service := services.NewQuoteSettingService(repository)

	result, err := service.CreateQuoteSetting(dto)

	if err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	c.JSON(http.StatusOK, successResponse(result))

}
