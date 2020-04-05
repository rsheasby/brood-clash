package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/rsheasby/brood-clash/backend/services"
	"github.com/rsheasby/brood-clash/backend/services/database"
)

func RevealAnswer (c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.Status(http.StatusBadRequest)
		return
	}
	err = database.RevealAnswer(id)
	if err != nil {
		c.Status(http.StatusNotFound)
		return
	}
	go services.BroadcastRevealAnswer(id)
	c.Status(http.StatusOK)
}
