package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rsheasby/brood-clash/backend/services"
)

func GetPresenterWebsocket(c *gin.Context) {
	err := services.Melody.HandleRequest(c.Writer, c.Request)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, fmt.Errorf("couldn't upgrade to websocket: %v", err))
		return
	}
}
