package models

import "github.com/go-playground/validator/v10"

type AppUser struct {
	ID   uint
	Name string `json:"name" validate:"required"`
	Age  int64  `json:"age" validate:"required"`
}

func AppUserStructValidation(supplied validator.StructLevel) {
	user := supplied.Current().Interface().(AppUser)
	if len(user.Name) == 0 {
		supplied.ReportError(user.Name, "name", "Name", "name", "")
	}
}
