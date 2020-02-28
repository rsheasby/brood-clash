package database

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/rsheasby/brood-clash/backend/models"
)

func RevealAnswer(ID uuid.UUID) error {
	a := new(models.Answer)
	err := db.Take(a, "id = ?", ID).Error
	if err != nil {
		return fmt.Errorf("error retrieving answer from database: %v", err)
	}
	a.Revealed = true
	db.Save(a)
	return nil
}
