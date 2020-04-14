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
		return ErrIDNotFound
	}
	if a.Revealed {
		return ErrAlreadyRevealed
	}
	a.Revealed = true
	db.Save(a)
	return nil
}

func GetAnswer(ID uuid.UUID) (answer *models.Answer, err error) {
	answer = new(models.Answer)

	err = db.Take(answer, "id = ?", ID).Error
	if err != nil {
		return nil, fmt.Errorf("error retrieving answer from database: %v", err)
	}
	return
}
