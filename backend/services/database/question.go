package database

import (
	"fmt"
	"github.com/rsheasby/brood-clash/backend/models"
)

func InsertQuestion(question models.Question) error {
	// Validate question
	if len(question.Answers) <= 0 {
		return fmt.Errorf("question must have at least one answer")
	} else if len(question.Answers) > 8 {
		return fmt.Errorf("question can't have more than eight answers")
	}

	// Validate answers
	var totalPoints int64
	for _, a := range question.Answers {
		if a.Points <= 0 {
			return fmt.Errorf("answers must have at least one point")
		}
		totalPoints += int64(a.Points)
	}
	if totalPoints > 100 {
		return fmt.Errorf("answer points can't total to over 100")
	}

	return fmt.Errorf("unable to insert question: %v", db.Create(&question).Error)
}

func GetUnshownQuestion() (result *models.Question, err error) {
	// Get unshown question
	result = new(models.Question)
	err = db.First(result, "has_been_shown = false").Error
	if err != nil {
		return nil, fmt.Errorf("unable to get unshown question: %v", err)
	}

	// Mark question as shown
	result.HasBeenShown = true
	err = db.Save(result).Error
	if err != nil {
		return nil, fmt.Errorf("unable to mark question as shown: %v", err)
	}

	return
}

func GetUnshownQuestionCount() (result int) {
	err := db.Model(&models.Question{}).Where("has_been_shown = false").Count(&result).Error
	if err != nil {
		return 0
	}
	return
}
