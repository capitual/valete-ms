package handlers

import (
	"net/http"
	"strconv"

	"github.com/capitual/valete_ms/internals/dtos"
	"github.com/capitual/valete_ms/internals/repositories"
	"github.com/capitual/valete_ms/internals/services"
	"github.com/gin-gonic/gin"
)

// NewPartner godoc
// @Summary      New Partner
// @Description  Create a new partner to consuming our api
// @Tags         partners
// @Accept       json
// @Produce      json
// @Param        partner body dtos.PartnerDto true "Add Partner"
// @Success      200  {object}  Response
// @Failure      400  {object}  HTTPError
// @Router       /internals/partners [post]
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

// GetPartner godoc
// @Summary      Get Partner
// @Description  get partner by id
// @Tags         partners
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "Partner id"
// @Param        x-api-secret  	header    string  true  "api secret"
// @Success      200  {object}  Response
// @Failure      400  {object}  HTTPError
// @Failure      401  {object}  HTTPUnauthorized
// @Router       /internals/partners/{id} [get]
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

// GetPartner godoc
// @Summary      Revogate Partner
// @Description  Revogate partner to use our api
// @Tags         partners
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "Partner id"
// @Param        x-api-secret  	header    string  true  "api secret"
// @Success      200  {object}  Response
// @Failure      400  {object}  HTTPError
// @Failure      401  {object}  HTTPUnauthorized
// @Router       /internals/partners/revogate/{id} [put]
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

// GetPartner godoc
// @Summary      Get All Partner
// @Description  Get all partner
// @Tags         partners
// @Accept       json
// @Produce      json
// @Param        page           query     int     false  "int valid"
// @Param        size           query     int     false  "int valid"
// @Param        filter         query     string  false  "string valid"
// @Param        x-api-secret  	header    string  true  "api secret"
// @Success      200  {object}  Response
// @Failure      400  {object}  HTTPError
// @Failure      401  {object}  HTTPUnauthorized
// @Router       /internals/partners [get]
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
