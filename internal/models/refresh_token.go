package models

import (
	"time"

	"gorm.io/gorm"
)

type RefreshToken struct {
	gorm.Model
	UserID    uint   `gorm:"not null"`
	User      User   `gorm:"foreignKey:UserID"`
	TokenHash string `gorm:"unique;not null"`
	ExpiresAt time.Time
}
