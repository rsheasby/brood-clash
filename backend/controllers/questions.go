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

// @Summary Get Unshown Question
// @ID get-unshown-question
// @Produce json
// @Security CodeAuth
// @Success 200 {object} models.Question "Success"
// @Failure 401 "Unauthorised"
// @Failure 404 "No more unshown questions"
// @Router /unshownQuestion [get]
func GetUnshownQuestion(c *gin.Context) {
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

// @Summary Get Current Question
// @ID get-current-question
// @Produce json
// @Security CodeAuth
// @Success 200 {object} models.Question "Success"
// @Failure 401 "Unauthorised"
// @Failure 404 "No current question"
// @Router /currentQuestion [get]
func GetCurrentQuestion(c *gin.Context) {
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

// @Summary Get All Questions
// @ID get-all-questions
// @Produce json
// @Security CodeAuth
// @Success 200 {array} models.Question "Success"
// @Failure 401 "Unauthorised"
// @Failure 500 "Failed to retrieve questions"
// @Router /questions [get]
func GetAllQuestions(c *gin.Context) {
	q, err := database.GetAllQuestions()
	if err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, q)
}

// @Summary Post Questions
// @ID post-questions
// @Accept json
// @Param questions body []models.Question true "Questions to be created"
// @Security CodeAuth
// @Success 201 "Questions created"
// @Failure 202 "Some questions couldn't be created"
// @Failure 401 "Unauthorised"
// @Router /questions [post]
func PostQuestions(c *gin.Context) {
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
		c.String(http.StatusAccepted, fmt.Sprintf("%d questions could not be created. %d have succeeded. Double check your entities are valid.", len(fails), len(questions)-len(fails)))
		return
	}
	c.Status(http.StatusCreated)
}
