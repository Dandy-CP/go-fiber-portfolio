package models

import "github.com/lib/pq"

type (
	MyProjects struct {
		Base
		ProjectName 		string					`json:"project_name" validate:"required,min=3"`
		Description 		string					`json:"description" validate:"required,min=3"`
		TechStack 			pq.StringArray	`json:"tech_stack" gorm:"type:text[];" validate:"required,min=1,dive"`
		DemoLink 				string					`json:"demo_link" validate:"required"`
		SourceCodeLink 	string					`json:"source_code_link" validate:"required"`
		ProjectImage 		string					`json:"project_image" validate:"required"`
	}
)