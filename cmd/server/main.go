package main

import (
	"github.com/MuhammadIbraAlfathar/online-store-app/config"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
)

func main() {

	godotenv.Load()
	config.LoadEnv()

	_, err := config.NewPostgres()
	if err != nil {
		log.Println("Failed connect to database")
	}

	r := gin.Default()
	r.GET("/ping", func(context *gin.Context) {
		context.JSON(200, gin.H{
			"message": "pong",
		})
	})

	err = r.Run()

	if err != nil {
		log.Println("Error run app")
	}
}
