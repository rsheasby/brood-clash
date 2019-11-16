package services

import (
	"backend/models"

	"github.com/davecgh/go-spew/spew"
)

var questions []models.Question

func CreateQuestion(q models.Question) {
	questions = append(questions, q)
}

func GetQuestions() []models.Question {
	return questions
}

func Spew() string {
	return spew.Sdump(questions)
}
