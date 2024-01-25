package main

import (
	"Ewallet/api"
	"Ewallet/db"

	"github.com/gin-gonic/gin"
)

func main() {
	// Инициализация маршрутов и подключение к базе данных
	r := gin.Default()
	db.InitDB()
	api.InitRoutes(r)

	// Запуск сервера
	r.Run(":8080")
}
