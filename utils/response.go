package utils

import (
	"github.com/gin-gonic/gin"
)

// ErrorResponse структура для ответа об ошибке
type ErrorResponse struct {
	Message string `json:"message"`
}

// RespondWithError отправляет ответ об ошибке клиенту
func RespondWithError(c *gin.Context, statusCode int, message string) {
	response := ErrorResponse{
		Message: message,
	}
	c.JSON(statusCode, response)
}
