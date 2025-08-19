package requests

import "smartjob/internal/models"

type QuestionRequest struct {
	Content string                  `json:"content"`
	Type    string                  `json:"type"`
	Options []models.QuestionOption `json:"options,omitempty"`
}
