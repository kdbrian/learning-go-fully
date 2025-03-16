package main

import (
	"errors"
	"fmt"
	"github.com/kdbrian/learning-go-fully/gorm_orm"
	. "github.com/kdbrian/learning-go-fully/internal/models"
	"gorm.io/gorm"
	"log"
)

func main() {

	if err := gorm_orm.ConnectSQLite(); err != nil {
		log.Fatal(err)
	}
	db := gorm_orm.SQLiteDB

	db.AutoMigrate(&User{})

	var user User
	println(errors.Is(db.First(&user, 1).Error, gorm.ErrRecordNotFound))

	var user_data map[string]interface{}
	db.Table("users").Take(&user_data)
	fmt.Printf("%+v\n", user_data)

	var appUser AppUser
	db.Model(&User{}).Find(&appUser)
	fmt.Printf("%+v\n", appUser)

}
