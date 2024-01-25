package db

import (
	"Ewallet/models"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

// InitDB инициализирует подключение к базе данных
func InitDB() {
	var err error
	dsn := "user=postgres password=root dbname=EWallet sslmode=disable"
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("Error connecting to database:", err)
		panic("Failed to connect to database")
	}

	// Отключение автоматических миграций (пересоздание таблиц)
	DB.AutoMigrate(&models.Wallet{}, &models.Transaction{})

}
