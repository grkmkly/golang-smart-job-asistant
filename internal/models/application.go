package models

import (
	"gorm.io/gorm"
)

type Application struct {
	gorm.Model         // CreatedAt = SubmittedAt
	Status     string  `gorm:"status;not null"`
	JobPostID  uint    `gorm:"index:idx_user_jobpost,unique"`
	JobPost    JobPost `gorm:"not null;foreignKey:JobPostID"`
	UserID     uint    `gorm:"index:idx_user_jobpost,unique"`
	User       User    `gorm:"foreignKey:UserID"`

	Answer []UserAnswer `gorm:"foreignKey:ApplicationID"`
}
