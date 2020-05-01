package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Summary Test
// @ID test
// @Security CodeAuth
// @Success 204 "Success"
// @Failure 401 "Unauthorised"
// @Router /test [get]
func Test(c *gin.Context) {
	c.Status(http.StatusNoContent)
}
