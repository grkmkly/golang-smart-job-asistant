package responses

import (
	"time"
)

type JobPostUserResponse struct {
	ID          uint      `json:"id"`
	Title       string    `json:"title"`
	Content     string    `json:"content"`
	EndAt       time.Time `json:"end_at"`
	IsActive    bool      `json:"is_active"`
	CreatedByID uint      `json:"created_by_id"`
}

type JobPostAdminResponse struct {
	Title        string                     `json:"title"`
	Content      string                     `json:"content"`
	EndAt        time.Time                  `json:"end_at"`
	IsActive     bool                       `json:"is_active"`
	CreatedByID  uint                       `json:"created_by_id"`
	JobQuestions []JobQuestionAdminResponse `json:"job_questions"`
}
