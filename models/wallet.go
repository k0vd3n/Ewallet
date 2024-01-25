package models

import "time"

// Wallet модель для кошелька
type Wallet struct {
	ID      string  `gorm:"type:varchar(32);primary_key" json:"id"`
	Balance float64 `gorm:"type:decimal(10,2);default:100.0" json:"balance"`
}

// Transaction модель для транзакции
type Transaction struct {
	ID         string    `gorm:"type:varchar(32);primary_key" json:"id"`
	WalletID   string    `gorm:"type:varchar(32);references:wallets;column:id" json:"wallet_id"`
	ToWalletID string    `gorm:"type:varchar(32);references:wallets;column:id" json:"to_wallet_id"`
	Amoung     float64   `gorm:"type:decimal(10,2)" json:"amount"`
	Time       time.Time `gorm:"type:timestamp without time zone;default:current_timestamp" json:"time"`
}
