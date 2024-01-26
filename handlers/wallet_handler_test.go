package handlers

import (
	mock_db "Ewallet/db/mocks"
	"Ewallet/models"
	"Ewallet/utils"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

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
