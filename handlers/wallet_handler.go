package handlers

import (
	"Ewallet/db"
	"Ewallet/models"
	"Ewallet/utils"
	"fmt"
	"net/http"
	"strings"
	"time"

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

// SendMoney обрабатывает запрос на перевод средств
func SendMoney(c *gin.Context) {
	// Получение идентификатора кошелька из параметра запроса
	walletID := c.Param("walletId")

	// Парсинг данных из тела запроса
	var request struct {
		To     string  `json:"to" binding:"required"`
		Amount float64 `json:"amount" binding:"required,min=0"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		// Возвращает ошибку в случае некорректных данных
		utils.RespondWithError(c, http.StatusBadRequest, "Invalid request data")
		return
	}
	// Проверка количества знаков после запятой в сумме перевода
	amountStr := fmt.Sprintf("%.2f", request.Amount)
	if strings.Contains(amountStr, ".") && len(amountStr[strings.Index(amountStr, ".")+1:]) > 2 {
		// Возвращает ошибку, если количество знаков после запятой больше двух
		utils.RespondWithError(c, http.StatusBadRequest, "Invalid amount format")
		return
	}

	// Начало транзакции
	tx := db.DB.Begin()

	// Обработка ошибок при начале транзакции
	if tx.Error != nil {
		utils.RespondWithError(c, http.StatusInternalServerError, "Failed to start transaction")
		return
	}

	defer func() {
		// В случае паники отменяем транзакцию
		if r := recover(); r != nil {
			tx.Rollback()
			utils.RespondWithError(c, http.StatusInternalServerError, "Transaction rolled back due to panic")
		}
	}()

	// Проверка, существует ли исходящий кошелек
	outgoingWallet := models.Wallet{}
	if err := tx.Where("id = ?", walletID).First(&outgoingWallet).Error; err != nil {
		// Возвращает ошибку, если кошелек не найден
		utils.RespondWithError(c, http.StatusNotFound, "Outgoing wallet not found")
		return
	}

	// Проверка, что исходящий кошелек не является кошельком, на который отправляются деньги
	if outgoingWallet.ID == request.To {
		// Возвращает ошибку, если исходящий кошелек совпадает с кошельком, на который отправляются деньги
		utils.RespondWithError(c, http.StatusBadRequest, "Cannot send money to the same wallet")
		return
	}

	// Создание новой транзакции
	newTransaction := models.Transaction{
		ID:         UuidWithoutHyphens(uuid.New()),
		WalletID:   outgoingWallet.ID,
		ToWalletID: UuidWithoutHyphens(uuid.MustParse(request.To)),
		Amount:     request.Amount,
		Time:       time.Now(),
	}

	// Вставка транзакции в базу данных
	if err := tx.Create(&newTransaction).Error; err != nil {
		// В случае ошибки вставки происходит отмена транзакции
		tx.Rollback()
		utils.RespondWithError(c, http.StatusBadRequest, "Failed to create transaction")
		return
	}

	// Обновление баланса исходяшего кошелька
	outgoingWallet.Balance -= request.Amount
	if err := tx.Save(&outgoingWallet).Error; err != nil {
		// В случае ошибки обновления баланса происходит отмента транзакции
		tx.Rollback()
		utils.RespondWithError(c, http.StatusBadRequest, "Failed to update outgoing wallet balance")
		return
	}

	// Получение информаци о входящем кошельке
	incomingWallet := models.Wallet{}
	if err := tx.Where("id = ?", request.To).First(&incomingWallet).Error; err != nil {
		// Возвращает ошибку, если входящий кошелек не был найден
		// Отмена транзакции
		tx.Rollback()
		utils.RespondWithError(c, http.StatusNotFound, "Incoming wallet not found")
		return
	}

	// Обновление баланса входящего кошелька
	incomingWallet.Balance += request.Amount
	if err := tx.Save(&incomingWallet).Error; err != nil {
		// В случае ошибки обеовления баланса происходит отмена транзакции
		tx.Rollback()
		utils.RespondWithError(c, http.StatusBadRequest, "Failed to update incoming wallet balance")
		return
	}

	// Фиксирование транзакции
	tx.Commit()

	// Возврат успешного ответа с информацией о транзакции
	c.JSON(http.StatusOK, newTransaction)
}

// GetTransactionHistory обрабатывает запрос на получение истории транзакций
func GetTransactionHistory(c *gin.Context) {
	walletID := c.Param("walletId")

	// Получение истории транзакций для указанного кошелька
	transactions := []models.Transaction{}
	if err := db.DB.Where("wallet_id = ? OR to_wallet_id = ?", walletID, walletID).Find(&transactions).Error; err != nil {
		// Возвращение ошибки, если не удалось получить историю транзакций
		utils.RespondWithError(c, http.StatusNotFound, "Wallet not found or error fetching transaction history")
		return
	}

	// Преобразуем историю транзакций в формат, соответствующий OpenAPI
	transactionHistory := make([]map[string]interface{}, 0)
	for _, transaction := range transactions {
		transactionEntry := map[string]interface{}{
			"time":   transaction.Time,
			"from":   transaction.WalletID,
			"to":     transaction.ToWalletID,
			"amount": transaction.Amount,
		}
		transactionHistory = append(transactionHistory, transactionEntry)
	}

	// Возвращаем успешный ответ с историей транзакций
	c.JSON(http.StatusOK, transactionHistory)
}

// GetWallet обрабатывает запрос на получение текущего состояния кошелька
func GetWallet(c *gin.Context) {
	walletID := c.Param("walletId")

	// Поиск кошелька в бд
	wallet := models.Wallet{}
	if err := db.DB.Where("id = ?", walletID).First(&wallet).Error; err != nil {
		// Если кошелек не найден возвращается ошибка 404
		utils.RespondWithError(c, http.StatusNotFound, "Wallet not found")
		return
	}

	// Возвращаем информацию о кошельке в успешном ответе
	c.JSON(http.StatusOK, wallet)
}

// UuidWithoutHyphens возвращает UUID в виде строки без дефисов
func UuidWithoutHyphens(u uuid.UUID) string {
	return strings.ReplaceAll(u.String(), "-", "")
}
