package models

import "gorm.io/gorm"

type Announcement struct {
	gorm.Model
	Title       string `json:"title" gorm:"size:255;not null"`
	Content     string `json:"content" gorm:"type:text;not null"`
	IsActive    bool   `json:"is_active" gorm:"default:true"`
	CreatedByID uint   `json:"createdBy_id" gorm:"not null"`
	CreatedBy   User   `gorm:"foreignKey:CreatedByID"`
}
