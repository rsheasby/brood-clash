package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rsheasby/brood-clash/backend/services/database"
)

// @Summary Reset Game State
// @ID reset-game-state
// @Security CodeAuth
// @Success 204 "Successfully reset the game state"
// @Failure 401 "Unauthorised"
// @Failure 500 "Failed to reset game state"
// @Router /reset [post]
func ResetGameState(c *gin.Context) {
	err := database.ResetGame()
	if err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	c.Status(http.StatusNoContent)
}
