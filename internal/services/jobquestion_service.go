package services

import (
	"smartjob/internal/mappers"
	"smartjob/internal/models"
	"smartjob/internal/requests"
	"smartjob/internal/responses"

	"gorm.io/gorm"
)

type JobQuestionService struct {
	DB *gorm.DB
}

func NewJobQuestionService(db *gorm.DB) *JobQuestionService {
	return &JobQuestionService{db}
}

func (s *JobQuestionService) Create(reqs []requests.JobQuestionRequest) error {
	questions, err := mappers.ReqsToJobQuestion(reqs)
	if err != nil {
		return err
	}
	if err = s.DB.Create(questions).Error; err != nil {
		return err
	}
	return nil
}
func (s *JobQuestionService) GetQuestionUserForPost(postID uint) ([]responses.JobQuestionUserResponse, error) {

	var jobQuestion []models.JobQuestion
	err := s.DB.Preload("Question").
		Preload("Question.Options").
		Where("job_post_id = ?", postID).
		Find(&jobQuestion).Error
	if err != nil {
		return nil, err
	}
	questionsResponse, err := mappers.JobQuestionsToUserResponse(jobQuestion)
	if err != nil {
		return nil, err
	}
	return questionsResponse, nil
}

func (s *JobQuestionService) GetQuestionAdminForPost(applicationID uint) ([]responses.JobQuestionAdminResponse, error) {
	var jobQuestion []models.JobQuestion

	err := s.DB.Preload("Question").
		Preload("Question.Options").
		Where("job_post_id = ?", applicationID).
		Find(&jobQuestion).Error
	if err != nil {
		return nil, err
	}
	questionsResponse, err := mappers.JobQuestionsToAdminResponse(jobQuestion)
	if err != nil {
		return nil, err
	}
	return questionsResponse, nil
}
