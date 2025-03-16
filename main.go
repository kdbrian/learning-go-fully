package main

import (
	"github.com/go-playground/validator/v10"
	"github.com/kdbrian/learning-go-fully/gorm_orm"
	. "github.com/kdbrian/learning-go-fully/internal/models"
	"log"
)

var validate *validator.Validate

func main() {

	validate = validator.New()

	//registering validator for appuser
	validate.RegisterStructValidation(AppUserStructValidation, AppUser{})

	if err := gorm_orm.ConnectSQLite(); err != nil {
		log.Fatal(err)
	}
	db := gorm_orm.SQLiteDB
	db.AutoMigrate(&User{})

	user := AppUser{
		Name: "",
		Age:  12,
	}
	errors := validate.Struct(&user)

	//errors := db.Create(&user).Error

	log.Fatal(errors)
}
