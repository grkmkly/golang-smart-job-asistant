package services

import (
	"smartjob/internal/mappers"
	"smartjob/internal/requests"

	"gorm.io/gorm"
)

type ApplicationService struct {
	DB *gorm.DB
}

func NewApplicationService(db *gorm.DB) *ApplicationService {
	return &ApplicationService{db}
}

func (s *ApplicationService) Create(req *requests.ApplicationRequest, userID uint) error {
	applications := mappers.ReqToApplicationModel(req, userID)

	if err := s.DB.Debug().Create(applications).Error; err != nil {
		return err
	}
	return nil
}

func (s *ApplicationService) GetApplicationsByPostID(postID uint) {
}
