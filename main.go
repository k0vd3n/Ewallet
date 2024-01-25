package main

import "github.com/gin-gonic/gin"

func main() {
	// Инициализация маршрутов и подключение к базе данных
	r := gin.Default()
	// Запуск сервера
	r.Run(":8080")
}
