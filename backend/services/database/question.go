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
	err := db.Create(&question).Error
	if err != nil {
		return fmt.Errorf("unable to insert question: %v", err)
	}
	return nil
}

func GetUnshownQuestion() (result *models.Question, err error) {
	// Get unshown question
	result = new(models.Question)
	err = db.First(result, "has_been_shown = false").Related(&result.Answers).Error
	if err != nil {
		return nil, fmt.Errorf("unable to get unshown question: %v", err)
	}

	tx := db.Begin()

	// Mark question as shown
	result.HasBeenShown = true
	err = tx.Save(result).Error
	if err != nil {
		return nil, fmt.Errorf("unable to mark question as shown: %v", err)
	}

	// Update current question in GameState
	err = tx.Model(&models.GameState{}).Update("question_id", result.ID).Error
	if err != nil {
		return nil, fmt.Errorf("unable to update GameState with new question ID: %v", err)
	}

	tx.Commit()

	return
}

func GetAllQuestions() (results []models.Question, err error) {
	err = db.Find(&results).Error
	return
}

func GetUnshownQuestionCount() (result int) {
	err := db.Model(&models.Question{}).Where("has_been_shown = false").Count(&result).Error
	if err != nil {
		return 0
	}
	return
}

func GetCurrentQuestionWithAnswers() (result *models.Question, err error) {
	gs := new(models.GameState)
	err = db.Take(gs).Error
	if err != nil {
		return nil, fmt.Errorf("unable to get current GameState: %v", err)
	}
	if gs.QuestionID == nil {
		return nil, ErrNoCurrentQuestion
	}
	result = new(models.Question)
	err = db.Take(result, "id = ?", gs.QuestionID).Related(&result.Answers).Error
	if err != nil {
		return nil, fmt.Errorf("unable to get current Question/Answers: %v", err)
	}
	return
}
