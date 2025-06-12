package controllers

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"warlock-backend/config"
	"warlock-backend/models"
	"warlock-backend/services"

	"github.com/gin-gonic/gin"
)

func SaveDocument() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		subjectIdStr := ctx.PostForm("subjectId")
		topicIdStr := ctx.PostForm("topicId")

		subjectId, err := strconv.Atoi(subjectIdStr)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid subject_id"})
			return
		}

		topicId, err := strconv.Atoi(topicIdStr)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid topic_id"})
			return
		}

		file, header, err := ctx.Request.FormFile("file")
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "File is required"})
			return
		}

		defer file.Close()

		basePath := "/app/filestore"
		subDirectory := fmt.Sprintf("/subject_%d/topic_%d", subjectId, topicId)
		fullPath := filepath.Join(basePath, subDirectory)

		mkdirErr := os.MkdirAll(fullPath, os.ModePerm)
		if mkdirErr != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create directory"})
			return
		}

		destDirectory := filepath.Join(fullPath, header.Filename)
		out, err := os.Create(destDirectory)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Could not create file"})
			return
		}
		defer out.Close()

		_, fileSaveError := out.ReadFrom(file)
		if fileSaveError != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save file"})
			return
		}

		document := models.Document{
			FileName:  header.Filename,
			FilePath:  destDirectory,
			SubjectId: uint(subjectId),
			TopicId:   uint(topicId),
		}
		if err := config.DB.Create(&document).Error; err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save file metadata"})
			return
		}

		log.Printf("Uploaded file: %s to %s", header.Filename, destDirectory)

		ctx.JSON(http.StatusOK, gin.H{
			"id":        document.ID,
			"file_name": document.FileName,
			"path":      document.FilePath,
		})

		// TODO: get userId from jwt
		services.NotifyAiService(document, 1, uint(subjectId), uint(topicId))

	}
}
