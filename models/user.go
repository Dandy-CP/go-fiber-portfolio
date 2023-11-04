package models

type (
	User struct {
		Base
		Username 	string 	`json:"username" validate:"required,min=3"`
		Password 	string 	`json:"password" validate:"required,min=8"`
	}
)