package database

import (
	"log"

	"github.com/rsheasby/brood-clash/backend/models"
)

func initGameState() {
	count := new(int)
	db.Model(&models.GameState{}).Count(count)
	switch {
	case *count == 0:
		db.Create(&models.GameState{QuestionID: nil})
	case *count == 1:
		break
	default:
		log.Fatalln("Corrupt database. Delete the database file and try again.")
	}
}

func ResetGame() (err error) {
	tx := db.Begin()

	err = tx.Model(models.GameState{}).Update("question_id", nil).Error
	if err != nil {
		tx.Rollback()
		return
	}

	err = tx.Model(models.Question{}).Update("has_been_shown", false).Error
	if err != nil {
		tx.Rollback()
		return
	}

	err = tx.Model(models.Answer{}).Update("revealed", false).Error
	if err != nil {
		tx.Rollback()
		return
	}

	tx.Commit()
	return nil
}
