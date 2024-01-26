package utils

import (
	"bytes"
	"encoding/json"
	"log"

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

// ParseResponseJSON разбирает JSON-ответ и помещает его в целевую структуру
func ParseResponseJSON(responseBody *bytes.Buffer, target interface{}) {
	err := json.NewDecoder(responseBody).Decode(target)
	if err != nil {
		log.Fatal(err)
	}
}
