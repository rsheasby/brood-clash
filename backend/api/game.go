package api

import (
	"backend/models"
	"backend/services"

	"net/http"

	"github.com/labstack/echo/v4"
)

func AddQuestions(c echo.Context) (err error) {
	var questions []models.Question
	if err = c.Bind(&questions); err != nil {
		return
	}

	for _, v := range questions {
		services.CreateQuestion(v)
	}

	return c.JSON(http.StatusOK, services.GetQuestions())
}

func GetQuestions(c echo.Context) (err error) {
	return c.JSON(http.StatusOK, services.GetQuestions())
}
