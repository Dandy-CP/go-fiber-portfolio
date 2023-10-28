package models

type (
	Blog struct {
		Base
		Thumbnail 	string 		`validate:"required"`
		Author 			string 		`validate:"required"`
		Title 			string 		`validate:"required"`
		Content 		string 		`validate:"required"`
	}
)