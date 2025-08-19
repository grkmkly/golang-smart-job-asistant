package models

import "gorm.io/gorm"

type Question struct {
	gorm.Model
	Content     string           `json:"content" gorm:"type:text;not null"`
	Type        string           `json:"type" gorm:"size:255;not null"`
	CreatedByID uint             `json:"created_By" gorm:"not null"`
	CreatedBy   User             `gorm:"foreignKey:CreatedByID"`
	Options     []QuestionOption `json:"options,omitempty" gorm:"foreignKey:QuestionID"`
}
type QuestionOption struct {
	gorm.Model
	QuestionID  uint     `json:"question_id" gorm:"not null"`
	Question    Question `gorm:"foreignKey:QuestionID"`
	OptionValue string   `json:"option_value" gorm:"size:255;not null"`
	Priority    int      `json:"priority"`
}
