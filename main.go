package main

import (
	"log"
	"time"
	"warlock-backend/config"
	"warlock-backend/controllers"
	"warlock-backend/cron"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	config.ConnectDB()
	go cron.CleanUpQaJob()

	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	router.GET("/health", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "Warlock backend is up and running",
		})
	})
	router.POST("/api/go/v1/signup", controllers.Signup())
	router.POST("/api/go/v1/login", controllers.Login())
	router.POST("/api/go/v1/validate-user", controllers.ValidateUser())
	router.PUT("/api/go/v1/add-strike-to-user", controllers.AddStrikeToUser())
	router.POST("/api/go/v1/save-qa", controllers.SaveQa())
	router.GET("/ws", func(c *gin.Context) {
		config.WsHandler(c.Writer, c.Request)
	})

	log.Println("Starting server on port 8080...")
	router.Run(":8080")
}
