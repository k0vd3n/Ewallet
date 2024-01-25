package api

import (
	"Ewallet/handlers"

	"github.com/gin-gonic/gin"
)

func InitRoutes(r *gin.Engine) {
	walletGroup := r.Group("/api/v1/wallet")
	{
		walletGroup.POST("/", handlers.CreateWallet)
		walletGroup.POST("/:walletId/send", handlers.SendMoney)
		walletGroup.GET("/:walletId/history", handlers.GetTransactionHistory)
		walletGroup.GET("/:walletId", handlers.GetWallet)
	}
}
