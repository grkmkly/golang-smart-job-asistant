package models

import (
	"gorm.io/gorm"
)

type JobQuestion struct {
	gorm.Model
	CriteriaValue    string `gorm:"type:varchar(255);column:criteria_value"`
	CriteriaOperator string `gorm:"type:varchar(50);column:criteria_operator"`
	QuestionOrder    int    `gorm:"column:question_order"`
	//Question
	QuestionID uint     `gorm:"not null"`
	Question   Question `gorm:"foreignKey:QuestionID"`

	JobPostID uint `gorm:"not null"`
}
