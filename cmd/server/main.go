package main

import (
	"github.com/MuhammadIbraAlfathar/online-store-app/config"
	"github.com/MuhammadIbraAlfathar/online-store-app/internal/domain/auth"
	"github.com/MuhammadIbraAlfathar/online-store-app/internal/domain/product"
	"github.com/MuhammadIbraAlfathar/online-store-app/internal/domain/user"
	"github.com/joho/godotenv"
	"log"
)

func main() {

	godotenv.Load()
	config.LoadEnv()

	engine := config.NewGin()

	db, err := config.NewPostgres()
	if err != nil {
		log.Println("Failed connect to database")
	}

	//USER
	userRepo := user.NewRepository(db)

	//AUTH
	authUseCase := auth.NewUseCase(userRepo)
	auth.NewController(engine, authUseCase)

	//PRODUCT
	productRepo := product.NewRepository(db)
	productUseCase := product.NewUseCase(productRepo)
	product.NewController(engine, productUseCase)

	err = engine.Run()
	if err != nil {
		log.Println("Error run app")
	}
}
