package models

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

type (
	Base struct {
		ID        uuid.UUID  `json:"id" gorm:"type:uuid;primary_key;"`
		CreatedAt time.Time  `json:"created_at"`
		UpdatedAt time.Time  `json:"update_at"`
		DeletedAt *time.Time `sql:"index" json:"deleted_at"`
	}

	ErrorResponse struct {
		Error       bool
		FailedField string
		Tag         string
		Value       interface{}
	}
)