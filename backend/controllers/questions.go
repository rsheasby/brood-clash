package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rsheasby/brood-clash/backend/models"
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
	c.JSON(http.StatusOK, q)
}

func GetCurrentQuestion (c *gin.Context) {
	q, err := database.GetCurrentQuestionWithAnswers()
	// TODO: Return message to client that there is no current question.
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
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
