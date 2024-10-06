package transaction

import (
	"github.com/MuhammadIbraAlfathar/online-store-app/internal/schema"
	"gorm.io/gorm"
)

type Repository interface {
	CreateTransaction(cart *schema.Transaction) (*schema.Transaction, error)
	GetTransactionById(transactionId int) (*schema.Transaction, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{db: db}
}

func (r *repository) CreateTransaction(transaction *schema.Transaction) (*schema.Transaction, error) {

	tx := r.db.Begin()

	err := tx.Create(&transaction).Error
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	tx.Commit()
	return transaction, nil
}

func (r *repository) GetTransactionById(transactionId int) (*schema.Transaction, error) {
	var transaction schema.Transaction
	err := r.db.Preload("Items.Product.Category").Preload("User").Where("id = ?", transactionId).First(&transaction).Error
	if err != nil {
		return nil, err
	}

	return &transaction, nil
}
