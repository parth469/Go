package database

import (
	models "github.com/parth469/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB = nil

func ConnectDatabase(dsn string) error {
	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}

	err = database.AutoMigrate(
		&models.User{},
	)
	if err != nil {
		return err
	}
	DB = database
	return nil
}

func GetDB() *gorm.DB {
	return DB
}
