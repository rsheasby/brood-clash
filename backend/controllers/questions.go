package controllers

import (
	"fmt"
	"net/http"

	"github.com/davecgh/go-spew/spew"
	"github.com/gin-gonic/gin"
	"github.com/rsheasby/brood-clash/backend/models"
	"github.com/rsheasby/brood-clash/backend/services/database"
)

func PostQuestions (c *gin.Context) {
	var questions []models.Question
	fails := make([]error, 0)
	err := c.BindJSON(&questions)
	if err != nil {
		c.Status(http.StatusBadRequest)
		return
	}
	spew.Dump(questions)
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
