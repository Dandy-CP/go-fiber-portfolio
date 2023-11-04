package models

type (
	Blog struct {
		Base
		Thumbnail 	string 		`json:"thumbnail" validate:"required"`
		Author 			string 		`json:"author" validate:"required"`
		Title 			string 		`json:"title" validate:"required"`
		Content 		string 		`json:"content" validate:"required"`
	}
)