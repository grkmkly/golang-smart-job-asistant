package responses

import (
	"time"
)

type QuestionUserResponse struct {
	ID      uint             `json:"id"`
	Content string           `json:"content"`
	Type    string           `json:"type"`
	Options []OptionResponse `json:"options"`
}
type OptionResponse struct {
	ID    int    `json:"id"`
	Value string `json:"value"`
}
type QuestionAdminResponse struct {
	ID        uint             `json:"id"`
	Content   string           `json:"content"`
	Type      string           `json:"type"`
	Options   []OptionResponse `json:"options"`
	CreatedBy UserResponse     `json:"created_by"`
	CreatedAt time.Time        `json:"created_at"`
}
