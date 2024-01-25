package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pedropezzuol/api_test/controller/request"
)

func CreateUser(c *gin.Context) {
	var userRequest request.UserRequest

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		c.JSON(http.StatusBadRequest, "error")
		return
	}

	c.JSON(200, userRequest)
}
