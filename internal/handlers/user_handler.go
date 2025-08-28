package handlers

import (
	"log"
	"net/http"
	"smartjob/internal/responses"
	"smartjob/internal/services"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	UserService *services.UserService
}

func NewUserHandler(userService *services.UserService) *UserHandler {
	return &UserHandler{
		UserService: userService,
	}
}

func (h *UserHandler) GetProfile() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		userID_interface, exists := ctx.Get("user_id")
		if !exists {
			responses.ErrorResponse(ctx, http.StatusInternalServerError, "GET_PROFILE_ERROR", &responses.APIError{
				Code:    "AUTHENCTICATION_ERROR",
				Details: "",
			})
		}
		userID, OK := userID_interface.(uint)
		if !OK {
			responses.ErrorResponse(ctx, http.StatusInternalServerError, "GET_PROFILE_ERROR", &responses.APIError{
				Code:    "INTERFACE_ERROR",
				Details: "",
			})
		}

		user, err := h.UserService.GetUserProfile(userID)
		if err != nil {
			log.Fatal("UserID Error", err)
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": err})
		}
		responses.SuccessResponse(ctx, http.StatusOK, "success", user)
	}
}
