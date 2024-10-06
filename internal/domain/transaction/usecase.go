package transaction

import (
	"errors"
	"github.com/MuhammadIbraAlfathar/online-store-app/internal/domain/cart"
	"github.com/MuhammadIbraAlfathar/online-store-app/internal/domain/product"
	"github.com/MuhammadIbraAlfathar/online-store-app/internal/schema"
	"gorm.io/gorm"
)

type UseCase struct {
	repo        Repository
	productRepo product.Repository
	cartRepo    cart.Repository
	db          *gorm.DB
}

func NewUseCase(repo Repository, productRepo product.Repository, cartRepo cart.Repository, db *gorm.DB) *UseCase {
	return &UseCase{
		repo:        repo,
		productRepo: productRepo,
		cartRepo:    cartRepo,
		db:          db,
	}
}

func (uc *UseCase) CreateTransaction(userId int) (*schema.Transaction, error) {
	carts, err := uc.cartRepo.GetCartByUserId(userId)
	if err != nil {
		return nil, errors.New("cart is empty")
	}

	tx := uc.db.Begin()
	if tx.Error != nil {
		return nil, tx.Error
	}

	var totalPrice int
	var transactionItems []schema.TransactionItem

	for _, item := range carts.CartItem {
		dataProduct := item.Product
		totalPrice += item.Quantity * dataProduct.Price
		transactionItems = append(transactionItems, schema.TransactionItem{
			ProductId: item.ProductId,
			Quantity:  item.Quantity,
			Price:     dataProduct.Price,
		})
	}

	transaction := schema.Transaction{
		UserId:     userId,
		TotalPrice: totalPrice,
		Items:      transactionItems,
	}

	newTransaction, err := uc.repo.CreateTransaction(&transaction)
	if err != nil {
		return nil, err
	}

	newDataTransaction, err := uc.repo.GetTransactionById(newTransaction.Id)
	if err != nil {
		return nil, err
	}

	return newDataTransaction, nil

}
