package models

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
)

type BaseModel struct {
	ID uuid.UUID `gorm:"type:string;primary_key;auto_increment:false" readonly:"true"`
}

func (base *BaseModel) BeforeCreate(scope *gorm.Scope) error {
	uuid, err := uuid.NewRandom()
	if err != nil {
		// I'm pretty sure this literally never errors, but I'm gonna put this here anyway since I can't prove it.
		uuidErr := fmt.Errorf("error generating UUID: %v", err)
		return scope.Err(uuidErr)
	}
	return scope.SetColumn("ID", uuid)
}
