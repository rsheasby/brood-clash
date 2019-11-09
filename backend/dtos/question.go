package dtos

import (
	"backend/models"
	"errors"
)

type Question struct {
	Description string   `json:"description"`
	Answers     []Answer `json:"answers"`
}

func (d *Question) ToModel() (model *models.Question, err error) {
	if len(d.Answers) > 8 {
		return nil, errors.New("Max 8 answers per question.")
	}

	model = new(models.Question)
	model.Description = d.Description
	for i := range d.Answers {
		a, _ := d.Answers[i].ToModel()
		model.Answers[i] = *a
	}

	return
}
