package controllers

import (
	"net/http"
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

		ctx.JSON(http.StatusOK, gin.H{"message": "User's number of strikes successfully implemented"})

	}
}
