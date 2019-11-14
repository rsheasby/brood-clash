package services

import (
	"backend/models"

	"github.com/davecgh/go-spew/spew"
)

var questions []models.Question

func init() {
	questions = make([]models.Question, 0)
}

func CreateQuestion(q models.Question) {
	questions = append(questions, q)
}

func GetQuestions() []models.Question {
	return questions
}

func Spew() string {
	return spew.Sdump(questions)
}
