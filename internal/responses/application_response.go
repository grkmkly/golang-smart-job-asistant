package responses

import "smartjob/internal/models"

type ApplicationResponse struct {
	ApplicationID uint                 `json:"application"`
	JobPost       models.JobPost       `json:"job_post"`
	Question      QuestionUserResponse `json:"question"`
	User          models.User          `json:"user_id"`
	Answer        UserAnswerResponse   `json:""`
}

type UserAnswerResponse struct {
}
