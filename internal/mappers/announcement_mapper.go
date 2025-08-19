package mappers

import (
	"smartjob/internal/models"
	"smartjob/internal/requests"
)

// REQ TO MODEL
func AnnouncementRequestToAnnouncement(req *requests.AnnouncementRequest, userID uint) *models.Announcement {
	return &models.Announcement{
		Title:       req.Title,
		Content:     req.Content,
		IsActive:    true,
		CreatedByID: userID,
	}
}
