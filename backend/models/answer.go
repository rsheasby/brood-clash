package models

import "github.com/google/uuid"

type Answer struct {
	BaseModel
	QuestionID uuid.UUID `gorm:"type:string;" json:"-"`
	Text       string    `json:"text,omitempty"`
	Points     int       `json:"points,omitempty" minimum:"1" maximum:"100"`
	Revealed   bool      `json:"revealed" readonly:"true"`
}
