package controllers

import (
	"log"
	"net/http"
	"warlock-backend/config"
	"warlock-backend/models"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func ValidateUserDocumentQa() gin.HandlerFunc {

	return func(ctx *gin.Context) {
		var input struct {
			WarlockApiKey string `json:"warlock_api_key"`
			SubjectId     uint   `json:"subject_id"`
			TopicId       uint   `json:"topic_id"`
		}
		claims := &models.Claims{}

		if err := ctx.ShouldBindJSON(&input); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
			return
		}

		token, err := jwt.ParseWithClaims(input.WarlockApiKey, claims, func(t *jwt.Token) (interface{}, error) {
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

		var subject models.Subject
		if err := config.DB.Preload("Topics").Where("id = ?", input.SubjectId).First(&subject).Error; err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "No subject found with the given subjectId"})
			return
		}

		isValidTopicForSubject := false
		for _, topic := range subject.Topics {
			if topic.ID == input.TopicId {
				isValidTopicForSubject = true
				log.Printf("Topic %v is related to Subject %v", topic.Name, subject.Name)
			}
		}

		if !isValidTopicForSubject {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "The topic is not related to the subject"})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"user_id":    user.ID,
			"user_roles": claims.Roles,
		})
	}
}
