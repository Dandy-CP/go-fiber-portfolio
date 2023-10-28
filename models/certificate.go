package models

type (
	Certificate struct {
		Base
		UrlPDF 					string		`validate:"required"`
		ThumbnailImage 	string		`validate:"required"`
	}
)