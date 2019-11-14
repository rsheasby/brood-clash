package api

import (
	"backend/dtos"
	"backend/models"

	"net/http"

	"github.com/davecgh/go-spew/spew"
	"github.com/labstack/echo/v4"
)

func AddQuestions(c echo.Context) (err error) {
	var q []dtos.Question
	if err = c.Bind(&q); err != nil {
		return
	}
	var models []models.Question
	for i := range q {
		m, _ := q[i].ToModel()
		models = append(models, *m)	
	}
	return c.String(http.StatusOK, spew.Sdump(q))
}
