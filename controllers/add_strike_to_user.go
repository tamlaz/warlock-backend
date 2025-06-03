package controllers

import (
	"log"
	"net/http"
	"os"
	"strconv"
	"warlock-backend/config"
	"warlock-backend/models"

	"github.com/gin-gonic/gin"
)

func AddStrikeToUser() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var input struct {
			UserId uint `json:"userId"`
		}

		if err := ctx.ShouldBindJSON(&input); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
			return
		}

		var user models.User
		if err := config.DB.Where("id = ?", input.UserId).First(&user).Error; err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{"error": "No user found with the given ID"})
			return
		}

		user.Strikes++
		if err := config.DB.Model(&user).Update("strikes", user.Strikes).Error; err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error":   "Could not update the number of striker due to error",
				"details": err.Error(),
			})
			return
		}

		maxStrikes, err := strconv.Atoi(os.Getenv("MAX_STRIKES"))
		if err != nil {
			log.Println("Failed to load max strikes from env")
			maxStrikes = 10
		}
		if user.Strikes == maxStrikes {
			if err := config.DB.Model(&user).Update("is_banned", true).Error; err != nil {
				ctx.JSON(http.StatusInternalServerError, gin.H{
					"error":   "Could not update isBanned flag on user",
					"details": err.Error(),
				})
				return
			}
			config.BroadcastToTopic("ban", gin.H{"userId": user.ID})
		}

		ctx.JSON(http.StatusOK, "User's number of strikes successfully implemented")

	}
}
