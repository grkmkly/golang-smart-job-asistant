package services

import (
	"smartjob/internal/mappers"
	"smartjob/internal/models"
	"smartjob/internal/requests"
	"smartjob/internal/responses"
	"time"

	"gorm.io/gorm"
)

type JobPostService struct {
	Db *gorm.DB
}

func NewJobPostService(db *gorm.DB) *JobPostService {
	return &JobPostService{db}
}
func (s *JobPostService) Create(req *requests.JobPostRequest, userID uint) error {
	newJobPost, err := mappers.JobPostRequestToJobPost(req, userID)
	if err != nil {
		return err
	}

	result := s.Db.Create(newJobPost)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
func (s *JobPostService) GetActiveAndNotExpiredPost() ([]responses.JobPostUserResponse, error) {
	var jobposts []models.JobPost

	now := time.Now()
	result := s.Db.
		Preload("JobQuestions").
		Preload("JobQuestions.Question").
		Preload("JobQuestions.Question.Options").
		Where("is_active = ?", true).
		Where("end_at > ?", now).
		Order("created_at DESC").
		Find(&jobposts)
	if result.Error != nil {
		return nil, result.Error
	}
	responses, err := mappers.JobPostModelToUserResponseSlice(&jobposts)
	if err != nil {
		return nil, err
	}
	return responses, nil
}

func (s *JobPostService) GetActiveAndNotExpiredPostForAdmin() ([]responses.JobPostAdminResponse, error) {
	var jobposts []models.JobPost

	now := time.Now()
	result := s.Db.
		Preload("JobQuestions").
		Preload("JobQuestions.Question").
		Preload("JobQuestions.Question.Options").
		Where("is_active = ?", true).
		Where("end_at > ?", now).
		Order("created_at DESC").
		Find(&jobposts)
	if result.Error != nil {
		return nil, result.Error
	}
	responses, err := mappers.JobPostModelToAdminResponseSlice(&jobposts)
	if err != nil {
		return nil, err
	}
	return responses, nil
}
