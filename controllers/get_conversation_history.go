package controllers

import (
	"log"
	"net/http"
	"strconv"
	"warlock-backend/config"
	"warlock-backend/models"

	"github.com/gin-gonic/gin"
)

func GetConversationHistory() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		userIdStr := ctx.Query("user_id")
		userId, err := strconv.ParseUint(userIdStr, 10, 64)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid or missing user_id query parameter"})
			return
		}

		if err := config.DB.Where("id = ?", userId).First(&models.User{}).Error; err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{"error": "No user found with the given ID"})
			return
		}

		var qaHistory []models.Qa
		if err := config.DB.Where("user_id = ?", userId).Find(&qaHistory).Order("created_at").Error; err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error fetching conversation history"})
			log.Printf("Error fetching conversation history: %v", err.Error())
			return
		}

		log.Println(qaHistory)

		lenQa := len(qaHistory)
		results := make([]models.HistoryMessage, lenQa*2)

		index := 0
		for _, qa := range qaHistory {
			historyMessageHuman := models.HistoryMessage{
				MessageContent: qa.Question,
				MessageType:    "HUMAN",
			}
			results[index] = historyMessageHuman
			index++
			historyMessageAi := models.HistoryMessage{
				MessageContent: qa.Answer,
				MessageType:    "AI",
			}
			results[index] = historyMessageAi
			index++
		}

		ctx.JSON(http.StatusOK, results)

	}
}
