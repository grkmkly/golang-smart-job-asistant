package mappers

import (
	"smartjob/internal/models"
	"smartjob/internal/requests"
)

// Request Announcement To Model Announcement
func AnnouncementRequestToAnnouncement(req *requests.AnnouncementRequest, userID uint) *models.Announcement {
	return &models.Announcement{
		Title:       req.Title,
		Content:     req.Content,
		IsActive:    true,
		CreatedByID: userID,
	}
}
