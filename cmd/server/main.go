package main

import (
	"github.com/MuhammadIbraAlfathar/online-store-app/config"
	"github.com/MuhammadIbraAlfathar/online-store-app/internal/domain/auth"
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

	//r := gin.Default()
	//r.GET("/ping", func(context *gin.Context) {
	//	context.JSON(200, gin.H{
	//		"message": "pong",
	//	})
	//})

	//USER
	userRepo := user.NewRepository(db)

	//AUTH
	authUseCase := auth.NewUseCase(userRepo)
	auth.NewController(engine, authUseCase)

	err = engine.Run()
	if err != nil {
		log.Println("Error run app")
	}
}
