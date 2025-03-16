package config

import (
	"fmt"
	"github.com/kdbrian/learning-go-fully/gorm_orm"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
)

var DB *gorm.DB

func dsn() string {
	return fmt.Sprintf(
		"host = %s port = %s user = %s password = %s dbname = %s sslmode = disable",
		os.Getenv("POSTGRES_DB_HOST"),
		os.Getenv("POSTGRES_DB_PORT"),
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_DB"),
	)
}

func Connect() error {
	db, err := gorm.Open(postgres.Open(dsn()), &gorm.Config{})
	if err != nil {
		return err
	}
	DB = db

	//if errStr := migrate(); errStr != "" {
	//	return errors.New(errStr)
	//}

	return nil
}

func migrate() {
	DB.AutoMigrate(&gorm_orm.AutoIncrementCounter{})
}
