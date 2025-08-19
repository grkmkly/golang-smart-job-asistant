package handlers

import (
	"net/http"
	"smartjob/internal/requests"
	"smartjob/internal/responses"
	"smartjob/internal/services"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	AuthService *services.AuthService
}

func NewAuthHandler(authService *services.AuthService) *AuthHandler {
	return &AuthHandler{
		AuthService: authService,
	}
}

func (h *AuthHandler) LoginUser() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		var req requests.LoginRequest

		if err := ctx.ShouldBindJSON(&req); err != nil {
			responses.ErrorResponse(ctx, http.StatusBadRequest, "LOGIN_FAILED", &responses.APIError{
				Code:    "LOGIN_BIND_FAILED",
				Details: err.Error(),
			})
			return
		}
		tokens, err := h.AuthService.LoginUser(req)
		if err != nil {
			responses.ErrorResponse(ctx, http.StatusBadRequest, "LOGIN_FAILED", &responses.APIError{
				Code:    "LOGIN_SERVICE_FAILED",
				Details: err.Error(),
			})
			return
		}
		responses.SuccessResponse(ctx, http.StatusOK, "LOGIN_SUCCESS", tokens)
	}
}

func (h *AuthHandler) RegisterUser() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		var req requests.RegisterRequest

		if err := ctx.ShouldBindJSON(&req); err != nil {
			responses.ErrorResponse(ctx, http.StatusBadRequest, "REGISTER_FAILED", &responses.APIError{
				Code:    "REGISTER_BIND_FAILED",
				Details: err.Error(),
			})
		}
		// RegisterUser
		user, err := h.AuthService.RegisterUser(req)
		if err != nil {
			responses.ErrorResponse(ctx, http.StatusBadRequest, "REGISTER_FAILED", &responses.APIError{
				Code:    "AUTH_SERVICE_FAILED",
				Details: err.Error(),
			})
			return
		}
		//Create User
		err = h.AuthService.UserService.CreateUser(&user)
		if err != nil {
			responses.ErrorResponse(ctx, http.StatusBadRequest, "REGISTER_FAILED", &responses.APIError{
				Code:    "USER_SERVICE_FAILED",
				Details: err.Error(),
			})
			return
		}
		responses.SuccessResponse(ctx, http.StatusOK, "LOGIN_SUCCESS", nil)
	}
}

func (h *AuthHandler) RefreshToken() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req requests.RefreshRequest
		if err := ctx.ShouldBindJSON(&req); err != nil {
			responses.ErrorResponse(ctx, http.StatusBadRequest, "REFRESH_TOKEN_FAILED", &responses.APIError{
				Code:    "REFRESH_BIND_FAILED",
				Details: err.Error(),
			})
			return
		}

		newAccessToken, err := h.AuthService.TokenService.RefreshAccessToken(req.RefreshToken)
		if err != nil {
			responses.ErrorResponse(ctx, http.StatusBadRequest, "REFRESH_TOKEN_FAILED", &responses.APIError{
				Code:    "TOKEN_SERVICE_FAILED",
				Details: err.Error(),
			})
			return
		}
		responses.SuccessResponse(ctx, http.StatusOK, "LOGIN_SUCCESS", gin.H{"access_token": newAccessToken})
	}

}
