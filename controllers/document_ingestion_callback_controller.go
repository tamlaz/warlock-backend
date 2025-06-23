package controllers

import (
	"net/http"
	"warlock-backend/config"

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

		message := gin.H{
			"topic":   "ingestion-success",
			"payload": gin.H{"documentId": input.DocumentId},
		}
		config.BroadcastToTopic("ingestion-success", message)

		ctx.JSON(http.StatusOK, gin.H{"status": "ok"})
	}
}
