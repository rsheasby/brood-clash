package controllers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/rotisserie/eris"
	"github.com/rsheasby/brood-clash/backend/models"
	"github.com/rsheasby/brood-clash/backend/services"
	"github.com/rsheasby/brood-clash/backend/services/database"
)

// @Summary Select Question
// @ID select-question
// @Produce json
// @Param id path string true "Question ID"
// @Security CodeAuth
// @Success 200 {object} models.Question "Success"
// @Failure 401 "Unauthorised"
// @Failure 404 "Question not found"
// @Failure 418 "Question has already been shown before"
// @Failure 500 "Unknown error"
// @Router /questions/{id}/select [post]
func SelectQuestion(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	question, err := database.SelectQuestion(id)
	if eris.Is(err, database.ErrIDNotFound) {
		c.Status(http.StatusNotFound)
		return
	}

	if eris.Is(err, database.ErrHasAlreadyBeenShown) {
		c.Status(http.StatusTeapot)
		return
	}

	if err != nil {
		c.Status(http.StatusInternalServerError)
	}

	go services.BroadcastStateUpdate()

	c.JSON(http.StatusOK, question)
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
		c.String(http.StatusNotFound, "No current question yet. You need to select a question first.")
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

// @Summary Delete Question
// @ID delete-question
// @Param id path string true "Question ID"
// @Security CodeAuth
// @Success 204 "Success"
// @Failure 400 "Couldn't parse ID param into UUID"
// @Failure 401 "Unauthorised"
// @Failure 404 "Question doesn't exist"
// @Failure 500 "Unknown error"
// @Router /questions/{id} [delete]
func DeleteQuestion(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	// In hindsight the deep deleting should've been implemented in a delete hook, but I can't be arsed to rewrite it now. If it comes up again, we can move it.
	err = database.DeleteQuestion(id)
	if eris.Is(err, database.ErrIDNotFound) {
		c.Status(http.StatusNotFound)
		return
	}
	if err != nil {
		log.Println(err)
		c.Status(http.StatusInternalServerError)
		return
	}

	c.Status(http.StatusNoContent)
}
