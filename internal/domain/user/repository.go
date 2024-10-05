package user

import (
	"errors"
	"github.com/MuhammadIbraAlfathar/online-store-app/internal/schema"
	"gorm.io/gorm"
)

type Repository interface {
	Create(user *schema.User) error
	FindById(id int) (*schema.User, error)
	FindByEmail(email string) (*schema.User, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{db: db}
}

func (r *repository) Create(user *schema.User) error {
	err := r.db.Create(&user).Error
	if err != nil {
		return err
	}

	return nil
}

func (r *repository) FindById(id int) (*schema.User, error) {
	var user schema.User
	if err := r.db.First(&user, id).Error; err != nil {
		return nil, errors.New("user not found")
	}

	return &user, nil
}

func (r *repository) FindByEmail(email string) (*schema.User, error) {
	var user schema.User
	if err := r.db.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, errors.New("user with that email was not found")
	}

	return &user, nil
}
