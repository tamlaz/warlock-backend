package main

import (
	"log"
	"warlock-backend/config"
	"warlock-backend/controllers"
	"warlock-backend/cron"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	config.ConnectDB()
	cron.CleanUpQaJob()

	router := gin.Default()

	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "Warlock backend is up and running",
		})
	})
	router.POST("/api/go/v1/signup", controllers.Signup())
	router.GET("/api/go/v1/login", controllers.Login())
	router.GET("/api/go/v1/validate-user", controllers.Login())

	router.Run(":8080")
}
