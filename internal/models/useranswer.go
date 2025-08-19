package models

import "gorm.io/gorm"

type UserAnswer struct {
	gorm.Model
	AnswerValue   string   `gorm:"not null;type:text"`
	ApplicationID uint     `gorm:"not null;index"`
	QuestionID    uint     `gorm:"not null"`
	Question      Question `gorm:"foreignKey:QuestionID"`
}
