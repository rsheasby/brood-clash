package database

import (
	"fmt"
	"sort"

	"github.com/google/uuid"
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

	//Sort answers
	sort.Slice(question.Answers, func(i, j int) bool {
		return question.Answers[i].Points > question.Answers[j].Points
	})

	err := db.Create(&question).Error
	if err != nil {
		return fmt.Errorf("unable to insert question: %v", err)
	}
	return nil
}

func SelectQuestion(id uuid.UUID) (result *models.Question, err error) {
	tx := db.Begin()

	result = new(models.Question)
	err = tx.Take(result, "id = ?", id).Error
	if err != nil {
		tx.Rollback()
		return nil, ErrIDNotFound
	}

	result.HasBeenShown = true
	err = tx.Save(result).Error
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	err = tx.Model(&models.GameState{}).Where("1 = 1").Update("question_id", result.ID).Error
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	tx.Commit()
	return
}

func GetAllQuestions() (results []models.Question, err error) {
	err = db.Find(&results).Error
	return
}

func DeleteQuestion(questionId uuid.UUID) (err error) {
	tx := db.Begin()

	q := new(models.Question)
	err = tx.Preload("Answers").Find(&q, "id = ?", questionId).Error
	if err != nil {
		tx.Rollback()
		return ErrIDNotFound
	}

	for i := range q.Answers {
		err = tx.Delete(q.Answers[i]).Error
		if err != nil {
			tx.Rollback()
			return err
		}
	}

	gs := new(models.GameState)
	err = tx.Take(gs).Error
	if err != nil {
		tx.Rollback()
		return err
	}

	if gs.QuestionID != nil && *gs.QuestionID == questionId {
		err = tx.Model(&models.GameState{}).Where("1 = 1").Update("question_id", nil).Error
		if err != nil {
			tx.Rollback()
			return err
		}
	}

	err = tx.Delete(&q).Error
	if err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
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
	err = db.Preload("Answers").Take(result, "id = ?", gs.QuestionID).Error
	if err != nil {
		return nil, fmt.Errorf("unable to get current Question/Answers: %v", err)
	}
	return
}
