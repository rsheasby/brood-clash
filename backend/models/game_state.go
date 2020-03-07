package models

import "github.com/google/uuid"

type GameState struct {
	QuestionID *uuid.UUID `gorm:"type:string;"`
}
