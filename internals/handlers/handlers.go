package handlers

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Message string      `json:"message,omitempty"`
	Body    interface{} `json:"data,omitempty"`
	Success bool        `json:"success"`
}

func getPagination(c *gin.Context) (int, int) {
	pageQ := c.Query("page")
	sizeQ := c.Query("size")

	var page int
	var size int

	var err error
	if page, err = strconv.Atoi(pageQ); err != nil {
		page = 1
	}

	if size, err = strconv.Atoi(sizeQ); err != nil {
		size = 10
	}

	return page, size
}

func errorResponse(err error) *Response {
	return &Response{
		Body:    nil,
		Success: false,
		Message: err.Error(),
	}
}

func successResponse(data interface{}) *Response {
	return &Response{
		Body:    data,
		Success: true,
	}
}
