package database

import (
	"github.com/rsheasby/brood-clash/backend/models"
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	_ "modernc.org/sqlite"
)

var db *gorm.DB

func init() {
	var err error // Yeah it's weird but it lets me directly set the global variable
	db, err = gorm.Open(sqlite.Dialector{
		DriverName: "sqlite",
		DSN:        "brood-clash.db",
	})
	if err != nil {
		log.Fatalln("Unable to open database: ", err)
	}

	db.AutoMigrate(
		models.Question{},
		models.Answer{},
		models.GameState{})

	initGameState()
}
