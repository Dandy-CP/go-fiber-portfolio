package models

type (
	Certificate struct {
		Base
		UrlPDF 					string		`json:"url_pdf" validate:"required"`
		ThumbnailImage 	string		`json:"thumbnail_image" validate:"required"`
	}
)