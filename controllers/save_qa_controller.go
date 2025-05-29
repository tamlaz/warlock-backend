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
			UserId   uint   `json:"userId"`
			Question string `json:"question"`
			Answer   string `json:"answer"`
		}

		if err := ctx.ShouldBindJSON(&input); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
			return
		}

		if err := config.DB.Where("id = ?", input.UserId).First(&models.User{}).Error; err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "User does not exist"})
			return
		}

		qa := models.Qa{
			UserId:   input.UserId,
			Question: input.Question,
			Answer:   input.Answer,
		}
		if err := config.DB.Create(&qa).Error; err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error during saving qa to DB"})
			return
		}

		ctx.JSON(http.StatusCreated, gin.H{"message": "New QA record saved successfully"})
	}
}
