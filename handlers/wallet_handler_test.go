package handlers

import (
	mock_db "Ewallet/db/mocks"
	"Ewallet/models"
	"Ewallet/utils"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	// Установка режима realese перед запуском всех тестов
	gin.SetMode(gin.ReleaseMode)

	// Запуск всех тестов и сохранение результатов покрытия в файл
	exitCode := m.Run()

	os.Exit(exitCode)
}

func TestCreateWallet(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// Создание мока для интерфейса WalletDatabase
	mockWalletDB := mock_db.NewMockWalletDatabase(ctrl)

	// Создание обработчика
	walletHandler := NewWalletHandler(mockWalletDB)

	// Создание фейкового Gin контекста
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	// Ожидаемый кошелек
	expectedWallet := &models.Wallet{
		ID:      "fake_id",
		Balance: 100.0,
	}

	// Устанавливаем ожидание вызова CreateWallet у mockWalletDB
	mockWalletDB.EXPECT().CreateWallet(gomock.Any()).Return(nil).SetArg(0, *expectedWallet)

	// Вызываем функцию CreateWallet
	walletHandler.CreateWallet(c)

	// Проверяем статус код и возвращенный кошелек
	assert.Equal(t, http.StatusOK, w.Code)

	var responseWallet models.Wallet
	utils.ParseResponseJSON(w.Body, &responseWallet)

	assert.Equal(t, expectedWallet.ID, responseWallet.ID)
	assert.Equal(t, expectedWallet.Balance, responseWallet.Balance)
}

func TestSendMoney(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// Создание макета интерфейса базы данных кошелька
	mockWalletDB := mock_db.NewMockWalletDatabase(ctrl)

	// Ожидаемый исходящий кошелек
	outgoingWallet := &models.Wallet{
		ID:      "9ff2881d0d7a4bd6831072f3af44c2f9",
		Balance: 200.0,
	}

	// Создание нового экземпляра обработчика кошелька
	walletHandler := NewWalletHandler(mockWalletDB)

	// Создание мнимого Gin контекста
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	// Установка параметра запроса, включая walletId
	c.Params = []gin.Param{{Key: "walletId", Value: outgoingWallet.ID}}

	// Ожидаемый входящий кошелек
	incomingWallet := &models.Wallet{
		ID:      "88b9af108d054efaae7a4c98105edffb",
		Balance: 50.0,
	}

	// Полезная нагрузка запроса
	requestPayload := fmt.Sprintf(`{"to": "%s", "amount": 50}`, incomingWallet.ID)

	// Установление ожидания для метода Begin
	mockWalletDB.EXPECT().Begin().Return(mockWalletDB).AnyTimes()

	// Установление ожидания для метода GetWalletByID
	mockWalletDB.EXPECT().GetWalletByID(outgoingWallet.ID).Return(outgoingWallet, nil)
	mockWalletDB.EXPECT().GetWalletByID(incomingWallet.ID).Return(incomingWallet, nil)

	// Установление ожидания для методов CreateTransaction и UpdateWalletBalance
	mockWalletDB.EXPECT().CreateTransaction(gomock.Any()).Return(nil).Do(func(transaction *models.Transaction) {
		assert.Equal(t, outgoingWallet.ID, transaction.WalletID)
		assert.Equal(t, incomingWallet.ID, transaction.ToWalletID)
		assert.Equal(t, 50.0, transaction.Amount)
		assert.True(t, transaction.Time.Before(time.Now().Add(time.Minute)))
	})
	mockWalletDB.EXPECT().UpdateWalletBalance(outgoingWallet).Return(nil)
	mockWalletDB.EXPECT().UpdateWalletBalance(incomingWallet).Return(nil)

	// Установление ожидания для метода Commit
	mockWalletDB.EXPECT().Commit().Return(nil)

	// Вызов функции SendMoney
	c.Request, _ = http.NewRequest(http.MethodPost, "/send-money/"+outgoingWallet.ID, strings.NewReader(requestPayload))
	walletHandler.SendMoney(c)

	// Проверка кода состояния ответа
	assert.Equal(t, http.StatusOK, w.Code)

	// Парсинг ответа в формате JSON
	var responseTransaction models.Transaction
	utils.ParseResponseJSON(w.Body, &responseTransaction)

	// Проверка информации об ответной транзакции
	assert.Equal(t, outgoingWallet.ID, responseTransaction.WalletID)
	assert.Equal(t, incomingWallet.ID, responseTransaction.ToWalletID)
	assert.Equal(t, 50.0, responseTransaction.Amount)
	assert.True(t, responseTransaction.Time.Before(time.Now().Add(time.Minute)))
}

func TestGetTransactionHistory(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// Создание мока для интерфейса WalletDatabase
	mockWalletDB := mock_db.NewMockWalletDatabase(ctrl)

	// Создание нового WalletHandler
	walletHandler := NewWalletHandler(mockWalletDB)

	// Создание фейкового Gin контекста
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	// Ожидаемый идентификатор кошелька
	expectedWalletID := "fake_wallet_id"

	// Установление параметры запроса
	c.Params = []gin.Param{{Key: "walletId", Value: expectedWalletID}}

	// Установление ожидания вызова GetWalletByID у mockWalletDB
	mockWalletDB.EXPECT().GetWalletByID(expectedWalletID).Return(nil, nil)

	// Ожидаемые транзакции
	expectedTransactions := []models.Transaction{
		{
			ID:         "fake_transaction_id_1",
			WalletID:   expectedWalletID,
			ToWalletID: "fake_to_wallet_id_1",
			Amount:     50.0,
			Time:       time.Now(),
		},
		{
			ID:         "fake_transaction_id_2",
			WalletID:   expectedWalletID,
			ToWalletID: "fake_to_wallet_id_2",
			Amount:     25.0,
			Time:       time.Now(),
		},
	}

	// Установление  ожидания вызова GetTransactionHistory у mockWalletDB
	mockWalletDB.EXPECT().GetTransactionHistory(expectedWalletID).Return(expectedTransactions, nil)

	// Вызов функции GetTransactionHistory
	walletHandler.GetTransactionHistory(c)

	// Проверка статуса кода и возвращенные транзакции
	assert.Equal(t, http.StatusOK, w.Code)

	var responseTransactions []map[string]interface{}
	utils.ParseResponseJSON(w.Body, &responseTransactions)

	// Проверка того, что количество возвращенных транзакций соответствует ожидаемому
	assert.Equal(t, len(expectedTransactions), len(responseTransactions))

	// Проверка деталей каждой транзакции
	for i, expectedTransaction := range expectedTransactions {
		assert.Equal(t, expectedTransaction.WalletID, responseTransactions[i]["from"])
		assert.Equal(t, expectedTransaction.ToWalletID, responseTransactions[i]["to"])
		assert.Equal(t, expectedTransaction.Amount, responseTransactions[i]["amount"])

		// Проверка того, что время соответствует ожидаемому с некоторой допустимой погрешностью
		expectedTimeStr := expectedTransactions[i].Time.Format(time.RFC3339Nano)
		assert.Equal(t, expectedTimeStr, responseTransactions[i]["time"])
	}
}

func TestGetWallet(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// Создание мока для интерфейса WalletDatabase
	mockWalletDB := mock_db.NewMockWalletDatabase(ctrl)

	// Создание обработчика
	walletHandler := NewWalletHandler(mockWalletDB)

	// Создание фейкового Gin контекста
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	// Ожидаемый идентификатор кошелька
	expectedWalletID := "fake_wallet_id"

	// Установление параметра запроса
	c.Params = []gin.Param{{Key: "walletId", Value: expectedWalletID}}

	// Ожидаемый кошелек
	expectedWallet := &models.Wallet{
		ID:      expectedWalletID,
		Balance: 100.0,
	}

	// Установление ожидания вызова GetWalletByID у mockWalletDB
	mockWalletDB.EXPECT().GetWalletByID(expectedWalletID).Return(expectedWallet, nil)

	// Вызов функции GetWallet
	walletHandler.GetWallet(c)

	// Проверка статуса кода и возвращенного кошелька
	assert.Equal(t, http.StatusOK, w.Code)

	var responseWallet models.Wallet
	utils.ParseResponseJSON(w.Body, &responseWallet)

	assert.Equal(t, expectedWallet.ID, responseWallet.ID)
	assert.Equal(t, expectedWallet.Balance, responseWallet.Balance)
}
