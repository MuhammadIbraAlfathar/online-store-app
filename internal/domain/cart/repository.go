package cart

import (
	"github.com/MuhammadIbraAlfathar/online-store-app/internal/schema"
	"gorm.io/gorm"
)

type Repository interface {
	CreateCart(cart *schema.Cart) (*schema.Cart, error)
	AddItem(cartId int, item *schema.CartItem) (*schema.CartItem, error)
	GetCartByUserId(userId int) (*schema.Cart, error)
	UpdateItem(cartId int, item *schema.CartItem) error
	ClearCart(cartID uint) error
	DeleteCartItem(req DeleteCartItemRequest) error
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{db: db}
}

func (r *repository) CreateCart(cart *schema.Cart) (*schema.Cart, error) {
	err := r.db.Create(&cart).Error
	if err != nil {
		return nil, err
	}

	return cart, nil
}

func (r *repository) GetCartByUserId(userId int) (*schema.Cart, error) {
	var cart schema.Cart
	err := r.db.Preload("CartItem.Product.Category").Where("user_id = ?", userId).First(&cart).Error
	if err != nil {
		return nil, err
	}

	return &cart, err
}

func (r *repository) AddItem(cartId int, item *schema.CartItem) (*schema.CartItem, error) {
	err := r.db.Create(&item).Error
	if err != nil {
		return nil, err
	}

	return item, nil
}

func (r *repository) UpdateItem(cartId int, item *schema.CartItem) error {
	return r.db.Save(item).Error
}

func (r *repository) ClearCart(cartID uint) error {
	return r.db.Where("cart_id = ?", cartID).Delete(&schema.CartItem{}).Error
}

func (r *repository) DeleteCartItem(req DeleteCartItemRequest) error {
	return r.db.Where("id = ?", req.ItemId).Delete(&schema.CartItem{}).Error
}
