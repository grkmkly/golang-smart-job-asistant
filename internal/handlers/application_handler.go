package handlers

import (
	"net/http"
	"smartjob/internal/requests"
	"smartjob/internal/responses"
	"smartjob/internal/services"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ApplicationHandler struct {
	ApplicationService services.ApplicationService
}

func NewApplicationHandler(AppS *services.ApplicationService) *ApplicationHandler {
	return &ApplicationHandler{*AppS}
}

func (h *ApplicationHandler) SubmitApplication() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req requests.ApplicationRequest

		if err := ctx.ShouldBindJSON(&req); err != nil {
			responses.ErrorResponse(ctx, http.StatusBadRequest, "SUBMIT_APPLICATION_ERROR", &responses.APIError{
				Code:    "BIND_ERROR",
				Details: err,
			})
			return
		}

		userID, _ := ctx.Get("user_id")
		postID := ctx.Param("jobpostID")
		postIDUint, err := strconv.ParseUint(postID, 10, 32)
		if err != nil {
			responses.ErrorResponse(ctx, http.StatusBadRequest, "SUBMIT_APPLICATION_ERROR", &responses.APIError{
				Code:    "PARSE_ERROR",
				Details: err,
			})
			return
		}

		err = h.ApplicationService.Create(&req, userID.(uint), uint(postIDUint))
		if err != nil {
			responses.ErrorResponse(ctx, http.StatusBadRequest, "SUBMIT_APPLICATION_ERROR", &responses.APIError{
				Code:    "SERVICE_ERROR",
				Details: err,
			})
			return
		}
		responses.SuccessResponse(ctx, http.StatusCreated, "success", nil)
	}
}

func (h *ApplicationHandler) GetApplicationsByPostID() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		postID := ctx.Param("jobpostID")
		postIDUint, err := strconv.ParseUint(postID, 10, 32)
		if err != nil {
			responses.ErrorResponse(ctx, http.StatusBadRequest, "GET_APPLICATIONS_ERROR", &responses.APIError{
				Code:    "PARSE_ERROR",
				Details: err,
			})
			return
		}
		response, err := h.ApplicationService.GetApplicationByPostIdWithCriteria((uint(postIDUint)))
		if err != nil {
			responses.ErrorResponse(ctx, http.StatusBadRequest, "GET_APPLICATIONS_ERROR", &responses.APIError{
				Code:    "SERVICE_ERROR",
				Details: err,
			})
			return
		}
		responses.SuccessResponse(ctx, http.StatusOK, "success", response)
	}
}
func (h *ApplicationHandler) UpdateStatus() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		appID := ctx.Param("applicationID")
		appIDUint, err := strconv.ParseUint(appID, 10, 32)
		if err != nil {
			responses.ErrorResponse(ctx, http.StatusBadRequest, "UPDATE_APPLICATIONS_ERROR", &responses.APIError{
				Code:    "PARSE_ERROR",
				Details: err,
			})
			return
		}
		var status requests.StatusRequest
		if err := ctx.ShouldBindJSON(&status); err != nil {
			responses.ErrorResponse(ctx, http.StatusBadRequest, "UPDATE_APPLICATIONS_ERROR", &responses.APIError{
				Code:    "BIND_ERROR",
				Details: err,
			})
			return
		}
		err = h.ApplicationService.UpdateApplicationStatus(uint(appIDUint), status.Status)
		if err != nil {
			responses.ErrorResponse(ctx, http.StatusBadRequest, "UPDATE_APPLICATIONS_ERROR", &responses.APIError{
				Code:    "SERVICE_ERROR",
				Details: err,
			})
			return
		}
		responses.SuccessResponse(ctx, http.StatusOK, "success", nil)
	}
}
func (h *ApplicationHandler) GetUserApplications() gin.HandlerFunc {
	return func(ctx *gin.Context) {

	}
}
