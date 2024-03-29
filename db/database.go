package db

import (
	"Ewallet/models"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// WalletDatabase интерфейс для базы данных
type WalletDatabase interface {
	CreateWallet(wallet *models.Wallet) error
	GetWalletByID(id string) (*models.Wallet, error)
	CreateTransaction(transaction *models.Transaction) error
	UpdateWalletBalance(wallet *models.Wallet) error
	GetTransactionHistory(walletID string) ([]models.Transaction, error)
	Begin() WalletDatabase
	Commit() error
	Rollback() error
	Error() error
}

/*
var (

	DB *gorm.DB

)
*/
type walletDatabase struct {
	*gorm.DB
}

// InitDB инициализирует подключение к базе данных и возвращает объект, удовлетворяющий интерфейсу WalletDatabase
func InitDB() WalletDatabase {

	// Загрузка переменных окружения из файла .env
	if err := godotenv.Load("db/.env"); err != nil {
		fmt.Println("Error loading .env file")
	}

	// Чтение значений переменных окружения
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", dbHost, dbPort, dbUser, dbPassword, dbName)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("Error connecting to database:", err)
		panic("Failed to connect to database")
	}

	db.AutoMigrate(&models.Wallet{}, &models.Transaction{})

	return &walletDatabase{DB: db}
}

func (w *walletDatabase) CreateWallet(wallet *models.Wallet) error {
	return w.DB.Create(wallet).Error
}

func (w *walletDatabase) GetWalletByID(id string) (*models.Wallet, error) {
	var wallet models.Wallet
	err := w.DB.Where("id = ?", id).First(&wallet).Error
	return &wallet, err
}

func (w *walletDatabase) CreateTransaction(transaction *models.Transaction) error {
	return w.DB.Create(transaction).Error
}

func (w *walletDatabase) UpdateWalletBalance(wallet *models.Wallet) error {
	return w.DB.Save(wallet).Error
}

func (w *walletDatabase) GetTransactionHistory(walletID string) ([]models.Transaction, error) {
	var transactions []models.Transaction
	err := w.DB.Where("wallet_id = ? OR to_wallet_id = ?", walletID, walletID).Find(&transactions).Error
	return transactions, err
}

func (w *walletDatabase) Begin() WalletDatabase {
	return &walletDatabase{DB: w.DB.Begin()}
}

func (w *walletDatabase) Commit() error {
	return w.DB.Commit().Error
}

func (w *walletDatabase) Rollback() error {
	return w.DB.Rollback().Error
}

func (w *walletDatabase) Error() error {
	return w.DB.Error
}
