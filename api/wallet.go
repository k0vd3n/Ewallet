package api

import (
	"Ewallet/db"
	"Ewallet/handlers"

	"github.com/gin-gonic/gin"
)

func InitRoutes(r *gin.Engine, walletDB db.WalletDatabase) {
	walletGroup := r.Group("/api/v1/wallet")
	{
		walletHandler := handlers.NewWalletHandler(walletDB)

		walletGroup.POST("/", walletHandler.CreateWallet)
		walletGroup.POST("/:walletId/send", walletHandler.SendMoney)
		walletGroup.GET("/:walletId/history", walletHandler.GetTransactionHistory)
		walletGroup.GET("/:walletId", walletHandler.GetWallet)
	}
}
