package handlers

import (
	"net/http"
	"smartjob/internal/requests"
	"smartjob/internal/responses"
	"smartjob/internal/services"
	"strconv"

	"github.com/gin-gonic/gin"
)

type JobQuestionHandler struct {
	JobQuestionService *services.JobQuestionService
}

func NewJobQuestioHandler(as *services.JobQuestionService) *JobQuestionHandler {
	return &JobQuestionHandler{as}
}

func (h *JobQuestionHandler) CreateJobQuestion() gin.HandlerFunc {

	return func(ctx *gin.Context) {
		var req []requests.JobQuestionRequest
		if err := ctx.ShouldBindJSON(&req); err != nil {
			responses.ErrorResponse(ctx, http.StatusBadRequest, "CREATE_JOB_QUESTION_ERROR", &responses.APIError{
				Code:    "BIND_ERROR",
				Details: err,
			})
			return
		}
		if err := h.JobQuestionService.Create(req); err != nil {
			responses.ErrorResponse(ctx, http.StatusBadRequest, "CREATE_JOB_QUESTION_ERROR", &responses.APIError{
				Code:    "SERVICE_ERROR",
				Details: err,
			})
			return
		}
		responses.SuccessResponse(ctx, http.StatusOK, "success", nil)
	}
}

func (h *JobQuestionHandler) GetQuestionUserForPost() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		jobpostIDStr := ctx.Param("jobpostID")
		postID, err := strconv.ParseUint(jobpostIDStr, 10, 32)
		if err != nil {
			responses.ErrorResponse(ctx, http.StatusInternalServerError, "GET_QUESTION_FOR_POST_ERROR", &responses.APIError{
				Code:    "PARSE_ERROR",
				Details: err,
			})
		}
		jobQuestions, err := h.JobQuestionService.GetQuestionUserForPost(uint(postID))

		if err != nil {
			responses.ErrorResponse(ctx, http.StatusInternalServerError, "GET_QUESTION_FOR_POST_ERROR", &responses.APIError{
				Code:    "SERVICE_ERROR",
				Details: err,
			})
			return
		}
		responses.SuccessResponse(ctx, http.StatusOK, "success", jobQuestions)
	}
}

func (h *JobQuestionHandler) GetQuestionAdminForPost() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		jobpostIDStr := ctx.Param("jobpostID")
		jobpostID, err := strconv.ParseUint(jobpostIDStr, 10, 32)
		if err != nil {
			responses.ErrorResponse(ctx, http.StatusInternalServerError, "GET_QUESTION_ADMIN_FOR_POST_ERROR", &responses.APIError{
				Code:    "PARSE_ERROR",
				Details: err,
			})
			return
		}
		jobQuestions, err := h.JobQuestionService.GetQuestionAdminForPost(uint(jobpostID))

		if err != nil {
			responses.ErrorResponse(ctx, http.StatusInternalServerError, "GET_QUESTION_ADMIN_FOR_POST_ERROR", &responses.APIError{
				Code:    "SERVICE_ERROR",
				Details: err,
			})
			return
		}
		responses.SuccessResponse(ctx, http.StatusOK, "success", jobQuestions)
	}
}
