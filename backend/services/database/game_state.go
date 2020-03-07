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
	default :
		log.Fatalln("Corrupt database. Delete the database file and try again.")
	}
}
