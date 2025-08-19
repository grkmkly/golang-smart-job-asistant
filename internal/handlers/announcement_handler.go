package handlers

import (
	"net/http"
	"smartjob/internal/requests"
	"smartjob/internal/responses"
	"smartjob/internal/services"

	"github.com/gin-gonic/gin"
)

type AnnouncementHandler struct {
	AnnouncementService *services.AnnouncementService
}

func NewAnnouncementHandler(as *services.AnnouncementService) *AnnouncementHandler {
	return &AnnouncementHandler{as}
}

func (h *AnnouncementHandler) CreateAnnouncement() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req requests.AnnouncementRequest
		if err := ctx.ShouldBindJSON(&req); err != nil {
			responses.ErrorResponse(ctx, http.StatusBadRequest, "CREATE_ANNOUNCEMENT_FAILED", &responses.APIError{
				Code:    "ANNOUNCEMENT_BIND_FAILED",
				Details: err.Error(),
			})
			return
		}
		userID, _ := ctx.Get("user_id")
		err := h.AnnouncementService.Create(&req, userID.(uint))
		if err != nil {
			responses.ErrorResponse(ctx, http.StatusBadRequest, "CREATE_ANNOUNCEMENT_FAILED", &responses.APIError{
				Code:    "ANNOUNCEMENT_SERVICE_FAILED",
				Details: err.Error(),
			})
			return
		}
		responses.SuccessResponse(ctx, http.StatusCreated, "success", nil)
	}
}

func (h *AnnouncementHandler) GetAnnouncements() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		announcements, err := h.AnnouncementService.GetAllActive()
		if err != nil {
			responses.ErrorResponse(ctx, http.StatusBadRequest, "GET_ANNOUNCEMENT_FAILED", &responses.APIError{
				Code:    "ANNOUNCEMENT_SERVICE_FAILED",
				Details: err.Error(),
			})
			return
		}
		responses.SuccessResponse(ctx, http.StatusOK, "success", announcements)
	}
}
