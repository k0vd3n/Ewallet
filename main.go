package main

import (
	"Ewallet/api"
	"Ewallet/db"

	"github.com/gin-gonic/gin"
)

func main() {
	// Инициализация подключения к базе данных
	walletDB := db.InitDB()
	// Создание роутера Gin
	r := gin.Default()
	// Инициализация маршрутов
	api.InitRoutes(r, walletDB)

	// Запуск сервера
	r.Run(":8080")
}
