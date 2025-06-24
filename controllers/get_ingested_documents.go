package controllers

import (
	"log"
	"net/http"
	"strconv"
	"warlock-backend/config"
	"warlock-backend/models"

	"github.com/gin-gonic/gin"
)

type IngestedDocument struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

func GetIngestedDocuments() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		userId, err := strconv.Atoi(ctx.Query("userId"))
		if err != nil {
			errorMessage := "Error fetching ingested documents"
			ctx.JSON(http.StatusInternalServerError, errorMessage)
			log.Printf("%v: %v", errorMessage, err.Error())
			return
		}
		var documents []models.Document
		if err := config.DB.Where("is_ingested = true AND user_id = ?", int(userId)).Find(&documents).Error; err != nil {
			ctx.JSON(http.StatusInternalServerError, "Error fetching ingested documents.")
			log.Printf("Error fetching ingested documents: %v", err.Error())
			return
		}

		ingestedDocuments := make([]IngestedDocument, len(documents))
		for i, document := range documents {
			ingestedDocuments[i] = IngestedDocument{
				ID:   document.ID,
				Name: document.FileName,
			}
		}

	}
}
