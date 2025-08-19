package mappers

import (
	"smartjob/internal/models"
	"smartjob/internal/requests"
)

// REQ TO MODEL
func JobPostRequestToJobPost(req *requests.JobPostRequest, userID uint) (*models.JobPost, error) {
	jobQuestions, err := ReqsToJobQuestion(req.JobQuestion)
	if err != nil {
		return nil, err
	}

	return &models.JobPost{
		Title:       req.Title,
		Content:     req.Content,
		EndAt:       req.EndAt,
		IsActive:    true,
		CreatedByID: userID,
		JobQuestion: jobQuestions,
	}, err
}
