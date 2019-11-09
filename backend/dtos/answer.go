package dtos

import "backend/models"

type Answer struct {
	Text   string `json:"text"`
	Points int    `json:"points"`
}

func (d *Answer) ToModel() (model *models.Answer, err error) {
	return &models.Answer{
		Text:   d.Text,
		Points: d.Points,
	}, nil
}
