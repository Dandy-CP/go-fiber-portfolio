package models

import "github.com/lib/pq"

type (
	MyProjects struct {
		Base
		ProjectName 		string					`validate:"required,min=3"`
		Description 		string					`validate:"required,min=3"`
		TechStack 			pq.StringArray	`gorm:"type:text[];" validate:"required,min=1,dive"`
		DemoLink 				string					`validate:"required"`
		SourceCodeLink 	string					`validate:"required"`
		ProjectImage 		string					`validate:"required"`
	}
)