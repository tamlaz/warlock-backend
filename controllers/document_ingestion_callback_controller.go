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

		config.BroadcastToTopic("ingestion-success", gin.H{"documentId": input.DocumentId})

		ctx.JSON(http.StatusOK, nil)
	}
}
