package api

import (
	"backend/dtos"
	"backend/models"

	"net/http"

	"github.com/davecgh/go-spew/spew"
	"github.com/labstack/echo/v4"
)

func AddQuestions(c echo.Context) (err error) {
	q := make([]dtos.Question, 0)
	if err = c.Bind(&q); err != nil {
		return
	}
	var models [8]models.Question
	for i := range q {
		m, _ := q[i].ToModel()
		models[i] = *m
	}
	return c.String(http.StatusOK, spew.Sdump(q))
}
