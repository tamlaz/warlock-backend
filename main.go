package main

import (
	"log"
	"time"
	"warlock-backend/config"
	"warlock-backend/controllers"
	"warlock-backend/cron"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	//err := godotenv.Load()
	//if err != nil {
	//	log.Fatal("Error loading .env file")
	//}
	log.Println("Starting application...")
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

	// health endpoint
	router.GET("/health", func(ctx *gin.Context) {
		ctx.JSON(200, "Warlock backend is up and running")
	})

	// ws handshake endpoint
	router.GET("/ws", func(c *gin.Context) {
		config.WsHandler(c.Writer, c.Request)
	})

	// REST API endpoints
	router.POST("/api/go/v1/signup", controllers.Signup())
	router.POST("/api/go/v1/login", controllers.Login())
	router.POST("/api/go/v1/validate-user", controllers.ValidateUser())
	router.POST("/api/go/v1/validate-user-document-qa", controllers.ValidateUserDocumentQa())
	router.POST("/api/go/v1/save-document", controllers.SaveDocument())
	router.POST("/api/go/v1/save-qa", controllers.SaveQa())
	router.POST("/api/go/v1/document-ingestion-callback", controllers.DocumentIngestionCallback())
	router.PUT("/api/go/v1/add-strike-to-user", controllers.AddStrikeToUser())
	router.GET("api/go/v1/get-conversation-history", controllers.GetConversationHistory())
	router.GET("api/go/v1/get-subjects", controllers.GetSubjects())
	router.GET("api/go/v1/get-topics", controllers.GetTopics())
	router.GET("api/go/v1/get-ingested-documents", controllers.GetIngestedDocuments())

	log.Println("Starting server on port 8080...")
	router.Run(":8080")
}
