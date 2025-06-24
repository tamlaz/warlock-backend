package controllers

import (
	"log"
	"net/http"
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

		var documents []models.Document
		if err := config.DB.Where("is_ingested = true").Find(&documents).Error; err != nil {
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

		ctx.JSON(http.StatusOK, ingestedDocuments)
	}
}
