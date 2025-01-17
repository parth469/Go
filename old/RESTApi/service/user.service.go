package service

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/parth469/database"
	"github.com/parth469/models"
	"gorm.io/gorm"
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

	isUserExist, error := FindUserByEmail(input.Email, db)
	if error != nil {
		return fmt.Errorf("error checking if user exists: %v", error)
	}

	if isUserExist {
		return fmt.Errorf("email id already register")
	}

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

func UpdateUser(updateDetail *models.UserUpdateInput) error {
	db := database.GetDB()

	isUserExist, err := FindUserByEmail(updateDetail.Email, db)
	if err != nil {
		return err
	}
	if !isUserExist {
		return fmt.Errorf("user does not exist with the given email")
	}

	hashedPassword, err := HashPassword(updateDetail.Password)
	if err != nil {
		return err
	}
	updateDetail.Password = hashedPassword

	if err := db.Model(&models.User{}).Where("email = ?", updateDetail.Email).Updates(&updateDetail).Error; err != nil {
		return err
	}
	return nil
}

func FindUserByEmail(email string, db *gorm.DB) (bool, error) {
	var user models.User
	result := db.Where("email = ?", email).First(&user)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return false, nil
	}
	if result.Error != nil {
		return false, result.Error
	}
	return true, nil
}
