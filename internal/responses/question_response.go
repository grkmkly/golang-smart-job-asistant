package responses

import (
	"smartjob/internal/models"
	"time"

	"gorm.io/gorm"
)

type QuestionUserResponse struct {
	gorm.Model
	Content string           `json:"content"`
	Type    string           `json:"type"`
	Options []OptionResponse `json:"options"`
}
type OptionResponse struct {
	ID    int    `json:"id"`
	Value string `json:"value"`
}
type QuestionAdminResponse struct {
	gorm.Model
	Content   string           `json:"content"`
	Type      string           `json:"type"`
	Options   []OptionResponse `json:"options"`
	CreatedBy models.User      `json:"created_by"`
	CreatedAt time.Time        `json:"created_at"`
}
