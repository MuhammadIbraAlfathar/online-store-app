package schema

import (
	"gorm.io/gorm"
	"time"
)

type Transaction struct {
	Id         int               `json:"id" gorm:"primaryKey;autoIncrement"`
	UserId     int               `json:"user_id" gorm:"not null"`
	User       User              `json:"user"`
	TotalPrice int               `json:"total_price" gorm:"not null"`
	Items      []TransactionItem `json:"items"`
	CreatedAt  time.Time         `json:"created_at"`
	UpdatedAt  time.Time         `json:"updated_at"`
	DeletedAt  gorm.DeletedAt    `gorm:"index"`
}

type TransactionItem struct {
	Id            int            `json:"id" gorm:"primaryKey;autoIncrement"`
	TransactionId int            `json:"transaction_id" gorm:"not null"`
	ProductId     int            `json:"product_id" gorm:"not null"`
	Quantity      int            `json:"quantity" gorm:"not null"`
	Price         int            `json:"price" gorm:"not null"`
	Product       Product        `json:"product"`
	CreatedAt     time.Time      `json:"created_at"`
	UpdatedAt     time.Time      `json:"updated_at"`
	DeletedAt     gorm.DeletedAt `gorm:"index"`
}
