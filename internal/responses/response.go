package responses

import "github.com/gin-gonic/gin"

type APIResponse struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
	Error   *APIError   `json:"error,omitempty"`
}

type APIError struct {
	Code    string      `json:"code,omitempty"`
	Details interface{} `json:"details,omitempty"`
}

func SuccessResponse(ctx *gin.Context, statusCode int, message string, data interface{}) {
	ctx.JSON(statusCode, APIResponse{
		Success: true,
		Message: message,
		Data:    data,
	})
}

func ErrorResponse(ctx *gin.Context, statusCode int, message string, errDetails *APIError) {
	ctx.JSON(statusCode, APIResponse{
		Success: false,
		Message: message,
		Error:   errDetails,
	})
}
