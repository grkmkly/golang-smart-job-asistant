package requests

import "time"

type JobPostRequest struct {
	Title       string               `json:"title"`
	Content     string               `json:"content"`
	EndAt       time.Time            `json:"end_at"`
	JobQuestion []JobQuestionRequest `json:"job_questions"`
}
