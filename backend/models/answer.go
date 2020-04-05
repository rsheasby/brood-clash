package models

import "github.com/google/uuid"

type Answer struct {
	BaseModel
	QuestionID uuid.UUID `gorm:"type:string;" json:"-"`
	Text string `json:",omitempty"`
	Points int `json:",omitempty"`
	Revealed bool
}
