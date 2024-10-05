package product

import (
	"github.com/MuhammadIbraAlfathar/online-store-app/internal/schema"
	"gorm.io/gorm"
)

type Repository interface {
	GetProductByCategoryId(categoryId int) (*[]schema.Product, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{db: db}
}

func (r *repository) GetProductByCategoryId(categoryId int) (*[]schema.Product, error) {
	var products []schema.Product
	if err := r.db.Preload("Category").Where("category_id = ?", categoryId).Find(&products).Error; err != nil {
		return nil, err
	}

	return &products, nil
}
