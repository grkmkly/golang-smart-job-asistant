package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	FirstName    string `json:"first_name" gorm:"size:255;not null"`
	Surname      string `json:"surname" gorm:"size:255;not null"`
	Email        string `json:"email" gorm:"size:255;unique;not null"`
	PasswordHash string `json:"-" gorm:"not null"`
	PhoneNumber  string `json:"phone_number" gorm:"size:20"`
	RoleID       uint   `json:"role_id" gorm:"not null;foreignKey:RoleID"`
	IsActive     bool   `json:"is_active" gorm:"default:true"`
}
