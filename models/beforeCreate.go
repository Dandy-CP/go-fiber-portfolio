package models

import (
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

func (value *Base) BeforeCreate(tx *gorm.DB) (err error) {
	value.ID = uuid.NewV4()
	return
}