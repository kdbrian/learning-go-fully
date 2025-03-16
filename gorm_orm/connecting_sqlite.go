package gorm_orm

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
)

var SQLiteDB *gorm.DB

func ConnectSQLite() error {
	db, err := gorm.Open(sqlite.Open("./.local/learning.db"))
	if err != nil {
		log.Fatal("failed : ", err)
		return err
	}

	SQLiteDB = db
	return nil
}
