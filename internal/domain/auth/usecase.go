package auth

import (
	"github.com/MuhammadIbraAlfathar/online-store-app/internal/domain/user"
	"github.com/MuhammadIbraAlfathar/online-store-app/internal/schema"
	"golang.org/x/crypto/bcrypt"
)

type UseCase struct {
	userRepo user.Repository
}

func NewUseCase(userRepo user.Repository) *UseCase {
	return &UseCase{
		userRepo: userRepo,
	}
}

func (uc *UseCase) Register(req *RegisterRequest) error {

	//Hash password
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	userEntity := schema.User{
		Name:     req.Name,
		Email:    req.Email,
		Password: string(passwordHash),
		Address:  req.Address,
	}

	if err := uc.userRepo.Create(&userEntity); err != nil {
		return err
	}

	return nil
}
