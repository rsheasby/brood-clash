package services

import (
	"backend/models"

	"github.com/davecgh/go-spew/spew"
)

var questions []models.Question

var nextID int = 1

func init() {
	questions = make([]models.Question, 0)
}

func CreateQuestion(q models.Question) (*models.Question, error) {
	if err := q.Validate(); err != nil {
		return nil, err
	}

	q.ID = nextID
	nextID += 1
	for i := range q.Answers {
		q.Answers[i].ID = i + 1
	}
	questions = append(questions, q)

	return GetQuestion(q.ID), nil
}

func GetQuestion(id int) *models.Question {
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
