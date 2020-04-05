package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rotisserie/eris"
	"github.com/rsheasby/brood-clash/backend/models"
	"github.com/rsheasby/brood-clash/backend/services"
	"github.com/rsheasby/brood-clash/backend/services/database"
)

func GetUnshownQuestion (c *gin.Context) {
	if database.GetUnshownQuestionCount() <= 0 {
		c.String(http.StatusNotFound, "No more questions.")
		return
	}
	q, err := database.GetUnshownQuestion()
	if err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}
	q, err = database.GetCurrentQuestionWithAnswers()
	if err != nil {
		c.Status(http.StatusInternalServerError)
	}
	go services.BroadcastStateUpdate()
	c.JSON(http.StatusOK, q)
}

func GetCurrentQuestion (c *gin.Context) {
	q, err := database.GetCurrentQuestionWithAnswers()
	if eris.Is(err, database.ErrNoCurrentQuestion) {
		c.String(http.StatusNotFound, "No current question yet. You should GET /unshownQuestion before calling me.")
		return
	}
	if err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, q)
}

func PostQuestions (c *gin.Context) {
	var questions []models.Question
	fails := make([]error, 0)
	err := c.BindJSON(&questions)
	if err != nil {
		c.Status(http.StatusBadRequest)
		return
	}
	for i := range questions {
		err := database.InsertQuestion(questions[i])
		if err != nil {
			c.Error(err)
			fails = append(fails, err)
		}
	}
	if len(fails) > 0 {
		c.String(http.StatusAccepted, fmt.Sprintf("%d questions could not be created. %d have succeeded. Double check your entities are valid.", len(fails), len(questions) - len(fails)))
		return
	}
	c.Status(http.StatusCreated)
}
