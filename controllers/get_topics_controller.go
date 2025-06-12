package controllers

import (
	"log"
	"net/http"
	"strconv"
	"warlock-backend/config"
	"warlock-backend/models"

	"github.com/gin-gonic/gin"
)

type TopicResult struct {
	ID        uint   `json:"id"`
	Name      string `json:"name"`
	SubjectId uint   `json:"subjectId"`
}

func GetTopics() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		subjectId, err := strconv.Atoi(ctx.Query("subjectId"))
		if err != nil {
			ctx.JSON(http.StatusBadRequest, "Required paramt subjectId is missing from request")
			log.Println("Error during retrieving subjectId: ", err.Error())
			return
		}

		var topics []models.Topic
		if err := config.DB.Where("subject_id = ?", subjectId).Find(&topics).Error; err != nil {
			ctx.JSON(http.StatusInternalServerError, "Cannot retrieve topics due to internal server error")
			log.Println("Error during fetching topics from db: ", err.Error())
			return
		}

		results := make([]TopicResult, len(topics))
		for i, subject := range topics {
			result := TopicResult{
				ID:        subject.ID,
				Name:      subject.Name,
				SubjectId: subject.SubjectId,
			}
			results[i] = result
		}

		ctx.JSON(http.StatusOK, results)
	}
}
