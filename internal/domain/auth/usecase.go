package auth

import (
	"errors"
	"github.com/MuhammadIbraAlfathar/online-store-app/internal/domain/user"
	"github.com/MuhammadIbraAlfathar/online-store-app/internal/schema"
	jwt_token "github.com/MuhammadIbraAlfathar/online-store-app/jwt"
	"golang.org/x/crypto/bcrypt"
	"log"
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

func (uc *UseCase) Login(req *LoginRequest) (*LoginResponse, error) {
	//get user by email
	dataUser, err := uc.userRepo.FindByEmail(req.Email)
	if err != nil {
		return nil, err
	}

	//compare hashing password
	err = bcrypt.CompareHashAndPassword([]byte(dataUser.Password), []byte(req.Password))
	if err != nil {
		return nil, errors.New("password not match")
	}

	accessToken, err := jwt_token.GenerateToken(dataUser.Id)
	if err != nil {
		log.Println("error access token")
		return nil, err
	}

	loginResponse := &LoginResponse{
		Id:    dataUser.Id,
		Name:  dataUser.Name,
		Email: dataUser.Email,
		Token: accessToken,
	}

	return loginResponse, nil
}
