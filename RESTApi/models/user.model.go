package models

type User struct {
	Base
	Name        string `json:"name" gorm:"type:varchar(100)"`
	PhoneNumber string `json:"phoneNumber" gorm:"type:varchar(15)"`
	Email       string `json:"email" gorm:"uniqueIndex;not null;type:varchar(100)"`
	Password    string `json:"-" gorm:"not null"`
	Status      bool   `json:"status" gorm:"default:true"`
}

type UserCreateInput struct {
	Name        string `json:"name" validate:"required,min=3,max=32"`
	PhoneNumber string `json:"phoneNumber" validate:"required,min=3,max=32"`
	Password    string `json:"password" validate:"required,min=6,max=8"`
	Email       string `json:"email" validate:"required,email"`
}
