package handlers

import (
	"net/http"
	"smartjob/internal/requests"
	"smartjob/internal/responses"
	"smartjob/internal/services"

	"github.com/gin-gonic/gin"
)

type QuestionHandler struct {
	QuestionService *services.QuestionService
}

func NewQuestionHandler(qs *services.QuestionService) *QuestionHandler {
	return &QuestionHandler{qs}
}

func (h *QuestionHandler) CreateQuestion() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req *requests.QuestionRequest
		if err := ctx.ShouldBindJSON(&req); err != nil {
			responses.ErrorResponse(ctx, http.StatusBadRequest, "CREATE_QUESTION_ERROR", &responses.APIError{
				Code:    "BIND_ERROR",
				Details: err,
			})
			return
		}
		userID, _ := ctx.Get("user_id")
		if err := h.QuestionService.CreateWithOption(req, userID.(uint)); err != nil {
			responses.ErrorResponse(ctx, http.StatusInternalServerError, "CREATE_QUESTION_ERROR", &responses.APIError{
				Code:    "QUESTION_SERVICE_ERROR",
				Details: err,
			})
			return
		}
		responses.SuccessResponse(ctx, http.StatusCreated, "success", nil)
	}
}

func (h *QuestionHandler) GetQuestionWithOption() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		questions, err := h.QuestionService.GetAdminWithOption()
		if err != nil {
			responses.ErrorResponse(ctx, http.StatusInternalServerError, "GET_QUESTION_ERROR", &responses.APIError{
				Code:    "QUESTION_SERVICE_ERROR",
				Details: err,
			})
		}
		responses.SuccessResponse(ctx, http.StatusOK, "success", questions)
	}
}
