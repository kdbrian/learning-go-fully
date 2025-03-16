package models

import (
	"errors"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name string
	Age  int
	Role string `gorm:"default:user"`
}

// Validate mostly deprecated
func (u User) Validate(db *gorm.DB) {

	//implement validation logic
	if len(u.Name) == 0 {
		db.AddError(errors.New("name is required"))
	}

}
