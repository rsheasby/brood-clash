package services

import (
	"backend/models"

	"github.com/davecgh/go-spew/spew"
)

// TODO: I think maybe this should be []*models.Question, not sure.
var questions []models.Question

var nextID int64 = 1

func init() {
	questions = make([]models.Question, 0)
}

func CreateQuestion(q models.Question) (*models.Question, error) {
	q.ID = nextID
	nextID += 1

	for i := range q.Answers {
		q.Answers[i].ID = int64(i + 1)
	}

	questions = append(questions, q)

	return GetQuestion(q.ID), nil
}

func GetQuestion(id int64) *models.Question {
	for _, v := range questions {
		if v.ID == id {
			return &v
		}
	}

	return nil
}

func GetQuestions() []models.Question {
	return questions
}

func Spew() string {
	return spew.Sdump(questions)
}
