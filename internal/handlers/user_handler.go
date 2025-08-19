package handlers

import (
	"log"
	"net/http"
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
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "No Authentication"})
		}
		userID, OK := userID_interface.(uint)
		if !OK {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "UserID is not valid"})
		}

		user, err := h.UserService.GetUserProfile(userID)
		if err != nil {
			log.Fatal("UserID Error", err)
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": err})
		}
		ctx.JSON(http.StatusOK, user)
	}
}
