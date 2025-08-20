package handlers

import (
	"net/http"
	"smartjob/internal/requests"
	"smartjob/internal/responses"
	"smartjob/internal/services"

	"github.com/gin-gonic/gin"
)

type JobPostHandler struct {
	JobPostService *services.JobPostService
}

func NewJobPostHandler(jbs *services.JobPostService) *JobPostHandler {
	return &JobPostHandler{jbs}
}

func (h *JobPostHandler) CreateNewJobPost() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req requests.JobPostRequest
		if err := ctx.ShouldBindJSON(&req); err != nil {
			responses.ErrorResponse(ctx, http.StatusBadRequest, "CREATE_NEW_JOB_ERROR", &responses.APIError{
				Code:    "BIND_ERROR",
				Details: err,
			})
			return
		}
		userID, _ := ctx.Get("user_id")
		err := h.JobPostService.Create(&req, userID.(uint))

		if err != nil {
			responses.ErrorResponse(ctx, http.StatusBadRequest, "CREATE_NEW_JOB_ERROR", &responses.APIError{
				Code:    "JOBPOST_SERVICE_ERROR",
				Details: err,
			})
			return
		}
		responses.SuccessResponse(ctx, http.StatusOK, "success", nil)
	}
}
func (h *JobPostHandler) ListJobPosts() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		jobposts, err := h.JobPostService.GetActiveAndNotExpiredPost()
		if err != nil {
			responses.ErrorResponse(ctx, http.StatusInternalServerError, "LIST_JOB_POSTS_ERROR", &responses.APIError{
				Code:    "JOB_POST_SERVICE_ERROR",
				Details: err,
			})
			return
		}
		responses.SuccessResponse(ctx, http.StatusOK, "success", jobposts)
	}
}

func (h *JobPostHandler) ListJobPostsForAdmin() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		jobposts, err := h.JobPostService.GetActiveAndNotExpiredPostForAdmin()
		if err != nil {
			responses.ErrorResponse(ctx, http.StatusInternalServerError, "LIST_JOB_POSTS_ERROR", &responses.APIError{
				Code:    "JOB_POST_SERVICE_ERROR",
				Details: err,
			})
			return
		}
		responses.SuccessResponse(ctx, http.StatusOK, "success", jobposts)
	}
}
