package handlers

import (
	"Ewallet/db"
	"Ewallet/models"
	"Ewallet/utils"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// CreateWallet обрабатывает запрос на создание кошелька
func CreateWallet(c *gin.Context) {
	// Генерация уникального ID для нового кошелька
	newUUID := uuid.New()
	newUUIDString := UuidWithoutHyphens(newUUID)

	newWallet := models.Wallet{
		ID:      newUUIDString,
		Balance: 100.0,
	}

	// Вставляем кошелек в базу данных
	if err := db.DB.Create(&newWallet).Error; err != nil {
		// В случае ошибки вставки, возвращаем ошибку
		utils.RespondWithError(c, http.StatusBadRequest, "Failed to create wallet")
		return
	}

	// Возвращаем успешный ответ с созданным кошельком
	c.JSON(http.StatusOK, newWallet)
}

// UuidWithoutHyphens возвращает UUID в виде строки без дефисов
func UuidWithoutHyphens(u uuid.UUID) string {
	return strings.ReplaceAll(u.String(), "-", "")
}
