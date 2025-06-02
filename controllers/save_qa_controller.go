package controllers

import (
	"net/http"
	"warlock-backend/config"
	"warlock-backend/models"

	"github.com/gin-gonic/gin"
)

func SaveQa() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var input struct {
			UserId              uint   `json:"user_id"`
			HumanMessageContent string `json:"human_message_content"`
			AiMessageContent    string `json:"ai_message_content"`
		}

		if err := ctx.ShouldBindJSON(&input); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
			return
		}

		var user models.User
		if err := config.DB.Where("id = ?", input.UserId).First(&user).Error; err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "User does not exist"})
			return
		}

		qa := models.Qa{
			User:     user,
			Question: input.HumanMessageContent,
			Answer:   input.AiMessageContent,
		}
		if err := config.DB.Create(&qa).Error; err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error during saving qa to DB"})
			return
		}

		ctx.JSON(http.StatusCreated, gin.H{"message": "New QA record saved successfully"})
	}
}
