package schema

import (
	"gorm.io/gorm"
	"time"
)

type Cart struct {
	Id        int            `json:"id" gorm:"primaryKey;autoIncrement"`
	UserId    int            `json:"user_id" gorm:"not null"`
	User      User           `json:"user"`
	CartItem  []CartItem     `json:"cart_item"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

type CartItem struct {
	Id        int            `json:"id" gorm:"primaryKey;autoIncrement"`
	CartId    int            `json:"cart_id" gorm:"not null"`
	ProductId int            `json:"product_id" gorm:"not null"`
	Product   Product        `json:"product"`
	Quantity  int            `json:"quantity" gorm:"not null"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
