package api

import (
	"backend/models"
	"backend/services"
	"fmt"

	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func AddQuestions(c echo.Context) (err error) {
	var questions []models.Question
	if err = c.Bind(&questions); err != nil {
		return
	}

	ids := make([]int, 0)
	for _, v := range questions {
		q, err := services.CreateQuestion(v)
		if err != nil {
			// TODO: Don't really know what to do with this error...
			fmt.Println(err)
		} else {
			ids = append(ids, q.ID)
		}
	}

	questions = make([]models.Question, 0)
	for _, id := range ids {
		questions = append(questions, *services.GetQuestion(id))
	}
	return c.JSON(http.StatusOK, questions)
}

func GetQuestion(c echo.Context) (err error) {
	id, _ := strconv.Atoi(c.Param("id"))
	question := services.GetQuestion(id)
	if question == nil {
		return c.NoContent(http.StatusNotFound)
	}
	return c.JSON(http.StatusOK, *question)

}

func GetQuestions(c echo.Context) (err error) {
	return c.JSON(http.StatusOK, services.GetQuestions())
}
