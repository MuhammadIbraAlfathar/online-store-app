package product

import "github.com/MuhammadIbraAlfathar/online-store-app/internal/schema"

type UseCase struct {
	repo Repository
}

func NewUseCase(repo Repository) *UseCase {
	return &UseCase{
		repo: repo,
	}
}

func (uc *UseCase) GetProductByCategoryId(categoryId int) (*[]schema.Product, error) {
	products, err := uc.repo.GetProductByCategoryId(categoryId)
	if err != nil {
		return nil, err
	}

	return products, err
}
