package models

import "github.com/google/uuid"

type Answer struct {
	BaseModel
	QuestionID uuid.UUID `gorm:"type:string;"`
	Text string
	Points int
	Revealed bool
}
