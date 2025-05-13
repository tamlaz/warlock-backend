package main

import (
	"log"
	"warlock-backend/config"
	"warlock-backend/controllers"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	config.ConnectDB()

	router := gin.Default()

	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "Warlock backend is up and running",
		})
	})
	router.POST("/signup", controllers.Signup())
	router.GET("/login", controllers.Login())

	router.Run(":8080")
}
