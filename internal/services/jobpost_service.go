package services

import (
	"smartjob/internal/mappers"
	"smartjob/internal/models"
	"smartjob/internal/requests"
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
func (s *JobPostService) GetActiveAndNotExpiredPost() ([]models.JobPost, error) {
	var jobposts []models.JobPost
	now := time.Now()
	result := s.Db.
		Preload("CreatedBy").
		Preload("JobQuestions").
		Where("is_active = ?", true).
		Where("end_at > ?", now).
		Order("created_at DESC").
		Find(&jobposts)

	if result.Error != nil {
		return nil, result.Error
	}
	return jobposts, nil
}
