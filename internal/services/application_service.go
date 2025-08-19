package services

import (
	"smartjob/internal/mappers"
	"smartjob/internal/models"
	"smartjob/internal/requests"
	"smartjob/internal/responses"

	"gorm.io/gorm"
)

type ApplicationService struct {
	DB                     *gorm.DB
	CriteriaControlService *CriteriaControlService
}

func NewApplicationService(db *gorm.DB, criteriaControlService *CriteriaControlService) *ApplicationService {
	return &ApplicationService{
		DB:                     db,
		CriteriaControlService: criteriaControlService,
	}
}

func (s *ApplicationService) Create(req *requests.ApplicationRequest, userID uint, postID uint) error {
	applications := mappers.ReqToApplicationModel(req, userID, postID)

	if err := s.DB.Debug().Create(applications).Error; err != nil {
		return err
	}
	return nil
}

func (s *ApplicationService) GetApplicationsByPostID(postID uint) ([]responses.ApplicationResponse, error) {
	var applications []models.Application

	result := s.DB.
		Preload("User").
		Preload("Answer").
		Preload("Answer.Question").
		Preload("JobPost").
		Preload("JobPost.JobQuestions").
		Where("job_post_id = ?", postID).
		Order("created_at DESC").
		Find(&applications)

	if result.Error != nil {
		return nil, result.Error
	}
	responses, err := mappers.ApplicationModelsToResponseSlice(applications)
	if err != nil {
		return nil, err
	}
	return responses, nil
}

func (s *ApplicationService) GetApplicationWithSuitable(applicationID uint) []responses.ResponseSuitable {
	application := models.Application{}
	if err := s.DB.First(&application, applicationID).Error; err != nil {
		return nil
	}
	return s.CriteriaControlService.CriteriaControl(application)
}

func (s *ApplicationService) GetApplicationByPostIdWithCriteria(postID uint) ([]responses.ResponseSuitable, error) {
	var applications []models.Application

	result := s.DB.
		Preload("User").
		Preload("Answer").
		Preload("Answer.Question").
		Preload("JobPost").
		Preload("JobPost.JobQuestions").
		Preload("JobPost.JobQuestions.Question").
		Where("job_post_id = ?", postID).
		Order("created_at DESC").
		Find(&applications)

	if result.Error != nil {
		return nil, result.Error
	}

	var suitableResponses []responses.ResponseSuitable
	for _, application := range applications {
		suitableResponses = append(suitableResponses, s.CriteriaControlService.CriteriaControl(application)...)
	}
	return suitableResponses, nil
}
