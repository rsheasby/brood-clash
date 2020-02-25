package controllers

import "github.com/gin-gonic/gin"

// ShowAccount godoc
// @Summary Ping the server to receive a pong
// @Success 200
// @Router /ping [get]
func Pong(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}
