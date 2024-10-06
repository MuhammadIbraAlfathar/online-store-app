package transaction

import (
	"time"
)

type ResponseTransaction struct {
	Id         int                     `json:"id"`
	UserId     int                     `json:"user_id"`
	User       UserTransactionResponse `json:"user"`
	TotalPrice int                     `json:"total_price"`
	Items      []ItemTransaction       `json:"items"`
	CreatedAt  time.Time               `json:"created_at"`
}

type UserTransactionResponse struct {
	Id      int    `json:"id"`
	Email   string `json:"email"`
	Address string `json:"address"`
}

type ItemTransaction struct {
	Id            int                 `json:"id"`
	TransactionId int                 `json:"cart_id"`
	ProductId     int                 `json:"product_id"`
	Quantity      int                 `json:"quantity"`
	Price         int                 `json:"price"`
	Product       ItemProductResponse `json:"product"`
}

type ItemProductResponse struct {
	Id           int    `json:"id"`
	Name         string `json:"name"`
	Price        int    `json:"price"`
	Stock        int    `json:"stock"`
	CategoryId   int    `json:"category_id"`
	CategoryName string `json:"category_name"`
}
