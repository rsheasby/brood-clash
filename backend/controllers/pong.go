package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/rsheasby/brood-clash/backend/services/database"
)

// @Summary Ping the server to receive a pong
// @Success 200
// @Router /ping [get]
func Pong(c *gin.Context) {
	count := database.GetUnshownQuestionCount()
	c.String(200, fmt.Sprint(count))
}
