package models

import (
	"time"

	"github.com/lib/pq"
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type Base struct {
	ID        uuid.UUID  `gorm:"type:uuid;primary_key;"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"update_at"`
	DeletedAt *time.Time `sql:"index" json:"deleted_at"`
}

type MyProjects struct {
	Base
	ProjectName 		string
	Description 		string
	TechStack 			pq.StringArray	`gorm:"type:text[]"`
	DemoLink 				string
	SourceCodeLink 	string
	ProjectImage 		string
}

func (value *Base) BeforeCreate(tx *gorm.DB) (err error) {
	value.ID = uuid.NewV4()
	return
}