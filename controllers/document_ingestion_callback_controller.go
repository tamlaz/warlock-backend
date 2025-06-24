package controllers

import (
	"log"
	"net/http"
	"warlock-backend/config"
	"warlock-backend/models"

	"fmt"

	"github.com/gin-gonic/gin"
)

func DocumentIngestionCallback() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var input struct {
			DocumentId uint `json:"document_id"`
		}

		if err := ctx.BindJSON(&input); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid payload"})
			return
		}

		var document models.Document
		if err := config.DB.Where("id = ?", input.DocumentId).First(&document).Error; err != nil {
			ctx.JSON(http.StatusInternalServerError, "No document found with given ID: "+fmt.Sprint(input.DocumentId))
			return
		}

		document.IsIngested = true
		if err := config.DB.Model(&document).Update("is_ingested", document.IsIngested).Error; err != nil {
			ctx.JSON(http.StatusInternalServerError, "Error during setting is_ingested on document with ID: "+fmt.Sprint(input.DocumentId))
			log.Printf("Error during setting is_ingested on document with ID %v: %v", input.DocumentId, err.Error())
			return
		}

		log.Printf("Document %v successfully ingested", document.FileName)

		message := gin.H{
			"topic":   "ingestion-success",
			"payload": gin.H{"documentId": input.DocumentId},
		}
		config.BroadcastToTopic("ingestion-success", message)

		ctx.JSON(http.StatusOK, gin.H{"status": "ok"})
	}
}
