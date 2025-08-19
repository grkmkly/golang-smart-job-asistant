package services

import (
	"smartjob/internal/mappers"
	"smartjob/internal/models"
	"smartjob/internal/requests"

	"gorm.io/gorm"
)

type AnnouncementService struct {
	DB *gorm.DB
}

func NewAnnouncementService(db *gorm.DB) *AnnouncementService {
	return &AnnouncementService{db}
}

func (s *AnnouncementService) Create(req *requests.AnnouncementRequest, UserID uint) error {
	announcement := mappers.AnnouncementRequestToAnnouncement(req, UserID)

	result := s.DB.Create(announcement)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
func (s *AnnouncementService) GetAllActive() ([]models.Announcement, error) {
	var announcements []models.Announcement
	result := s.DB.Preload("CreatedBy").Where("is_active = ?", true).Order("created_at DESC").Find(&announcements)
	if result.Error != nil {
		return nil, result.Error
	}
	return announcements, nil
}
