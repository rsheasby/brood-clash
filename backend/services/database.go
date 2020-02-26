package services

import (
	"fmt"
	"github.com/rsheasby/brood-clash/backend/models"
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var db *gorm.DB

func init() {
	var err error // Yeah it's weird but it lets me directly set the global variable
	db, err = gorm.Open("sqlite3", "brood-clash.db")
	if err != nil {
		log.Fatalln("Unable to open database: ", err)
	}

	db.AutoMigrate(
		models.Question{},
		models.Answer{})
}

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

	db.Create(&question)
	return nil
}
