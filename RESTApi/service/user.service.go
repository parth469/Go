package service

import (
	"encoding/json"
	"fmt"

	"github.com/parth469/database"
	"github.com/parth469/models"
)

type usersType []models.User

func AllUserFetch() (usersType, error) {
	db := database.GetDB()
	users := usersType{}

	result := db.Find(&users)

	if result.Error != nil {
		return nil, result.Error
	}

	_, err := json.Marshal(users)
	fmt.Println(users)
	if err != nil {
		return nil, err

	}
	return users, nil
}

func CreateUser(input models.UserCreateInput) error {
	db := database.GetDB()
	hashpass, err := HashPassword(input.Password)

	if err != nil {
		return err
	}

	user := models.User{Name: input.Name, Password: hashpass, PhoneNumber: input.PhoneNumber, Email: input.Email, Status: true}

	if err := db.Create(&user).Error; err != nil {
		return fmt.Errorf("failed to create user: %v", err)
	}

	return nil
}
