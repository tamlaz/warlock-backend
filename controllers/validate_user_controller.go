package controllers

import (
	"net/http"
	"warlock-backend/config"
	"warlock-backend/models"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func ValidateUser() gin.HandlerFunc {

	return func(ctx *gin.Context) {
		var input struct {
			Token string `json:"token"`
		}
		claims := &models.Claims{}

		token, err := jwt.ParseWithClaims(input.Token, claims, func(t *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})

		if err != nil || !token.Valid {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired token"})
			return
		}

		var user models.User
		if err := config.DB.Where("email = ?", claims.Email).First(&user).Error; err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "No user found with the given email"})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"userId": user.ID,
			"roles":  claims.Roles,
		})
	}
}
