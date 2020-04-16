package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/rotisserie/eris"
	"github.com/rsheasby/brood-clash/backend/services"
	"github.com/rsheasby/brood-clash/backend/services/database"
)

// @Summary Reveal answer
// @ID reveal-answer
// @Accept json
// @Produce json
// @Param id path string true "Answer ID, must be UUID"
// @Security CodeAuth
// @Success 204 "Success"
// @Failure 400 "Couldn't parse ID param into UUID"
// @Failure 403 "Unauthorised"
// @Failure 404 "Answer doesn't exist"
// @Failure 418 {string} string "Answer already revealed"
// @Router /answers/{id}/reveal [post]
func RevealAnswer(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.Status(http.StatusBadRequest)
		return
	}
	err = database.RevealAnswer(id)
	if eris.Is(err, database.ErrAlreadyRevealed) {
		c.String(http.StatusTeapot, "Answer is already revealed.")
		return
	}
	if err != nil {
		c.Status(http.StatusNotFound)
		return
	}
	go services.BroadcastRevealAnswer(id)
	c.Status(http.StatusNoContent)
}

// @Summary Incorrect Answer
// @ID incorrect-answer
// @Security CodeAuth
// @Success 204 "Success"
// @Failure 401 "Unauthorised"
// @Router /incorrectAnswer [post]
func IncorrectAnswer(c *gin.Context) {
	go services.BroadcastIncorrectAnswer()
	c.Status(http.StatusNoContent)
}
