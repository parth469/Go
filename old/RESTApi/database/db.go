package database

import (
	"fmt"

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
	sqlDB, err := database.DB()
	if err != nil {
		fmt.Println(err)
	}
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(0)
	err = database.AutoMigrate(
		&models.User{},
	)
	fmt.Println(err)
	if err != nil {
		return err
	}
	DB = database
	return nil
}

func GetDB() *gorm.DB {
	return DB
}
