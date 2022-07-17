package handlers

import (
	"net/http"
	"strconv"

	"github.com/capitual/valete_ms/internals/dtos"
	"github.com/capitual/valete_ms/internals/repositories"
	"github.com/capitual/valete_ms/internals/services"
	"github.com/gin-gonic/gin"
)

func NewPartner(c *gin.Context) {
	var p dtos.PartnerDto

	err := c.ShouldBindJSON(&p)

	if err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	repository := &repositories.PartnersRepository{}
	service := services.NewPartnerService(repository)

	result, err := service.CreatePartner(p)

	if err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	c.JSON(http.StatusOK, successResponse(result))
}

func GetPartnerById(c *gin.Context) {
	id := c.Param("id")

	newid, err := strconv.Atoi(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	repository := &repositories.PartnersRepository{}
	service := services.NewPartnerService(repository)

	partner, err := service.GetPartnerById(newid)

	if err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	c.JSON(http.StatusOK, successResponse(partner))
}

func RevogatePartner(c *gin.Context) {
	id := c.Param("id")

	newid, err := strconv.Atoi(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	repository := &repositories.PartnersRepository{}
	service := services.NewPartnerService(repository)

	success, err := service.RevogatePartner(newid)

	if err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	c.JSON(http.StatusOK, successResponse(success))
}

func GetAllPartners(c *gin.Context) {
	page, size, filter := getPagination(c)

	repository := &repositories.PartnersRepository{}
	service := services.NewPartnerService(repository)

	result, err := service.GetAll(filter, page, size)

	if err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	c.JSON(http.StatusOK, successResponse(result))
}
