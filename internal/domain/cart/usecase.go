package cart

import (
	"errors"
	"github.com/MuhammadIbraAlfathar/online-store-app/internal/domain/product"
	"github.com/MuhammadIbraAlfathar/online-store-app/internal/schema"
	"gorm.io/gorm"
)

type UseCase struct {
	repo        Repository
	productRepo product.Repository
}

func NewUseCase(repo Repository, productRepo product.Repository) *UseCase {
	return &UseCase{
		repo:        repo,
		productRepo: productRepo,
	}
}

func (uc *UseCase) AddProductToCart(req []AddToCartItemsRequest, userId int) (*[]schema.CartItem, error) {
	var cartItemData []schema.CartItem

	//check if the user has created a cart
	cart, err := uc.repo.GetCartByUserId(userId)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			cart = &schema.Cart{
				UserId: userId,
			}

			_, err := uc.repo.CreateCart(cart)
			if err != nil {
				return nil, err
			}
		}
	}

	for _, p := range req {
		dataProduct, err := uc.productRepo.GetProductById(p.ProductId)
		if err != nil {
			return nil, err
		}

		if dataProduct.Stock < p.Quantity {
			return nil, errors.New("no stock")
		}

		newItem := &schema.CartItem{
			CartId:    cart.Id,
			ProductId: p.ProductId,
			Quantity:  p.Quantity,
		}

		cartItem, err := uc.repo.AddItem(cart.Id, newItem)
		if err != nil {
			return nil, err
		}

		cartItemData = append(cartItemData, *cartItem)
	}

	_, err = uc.repo.GetCartByUserId(userId)
	if err != nil {
		return nil, err
	}

	return &cartItemData, err
}

func (uc *UseCase) GetCartUserByUserId(userId int) (*schema.Cart, error) {
	cart, err := uc.repo.GetCartByUserId(userId)
	if err != nil {
		return nil, err
	}

	return cart, nil
}
