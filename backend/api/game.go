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

func RevealAnswer(c echo.Context) (err error) {
	questionId, err := strconv.Atoi(c.Param("questionId"))
	if err != nil {
		fmt.Println(err)
		return c.NoContent(http.StatusBadRequest)
	}
	question := services.GetQuestion(questionId)
	if question == nil {
		return c.NoContent(http.StatusNotFound)
	}

	answerId, err := strconv.Atoi(c.Param("answerId"))
	if err != nil {
		fmt.Println(err)
		return c.NoContent(http.StatusBadRequest)
	}
	for i := range question.Answers {
		if question.Answers[i].ID == answerId {
			question.Answers[i].Revealed = true
			return c.NoContent(http.StatusOK)
		}
	}
	return c.NoContent(http.StatusNotFound)
}
