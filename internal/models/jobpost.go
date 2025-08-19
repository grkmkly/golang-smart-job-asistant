package models

import (
	"time"

	"gorm.io/gorm"
)

type JobPost struct {
	gorm.Model
	Title        string        `gorm:"size:255;not null"`
	Content      string        `gorm:"type:text;not null"`
	EndAt        time.Time     `gorm:"index"`
	IsActive     bool          `gorm:"default:true"`
	CreatedByID  uint          `gorm:"not null"`
	CreatedBy    User          `gorm:"foreignKey:CreatedByID"`
	JobQuestions []JobQuestion `gorm:"foreignKey:JobPostID"`
}
