package database

import (
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

