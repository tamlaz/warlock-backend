package controllers

import (
	"log"
	"net/http"
	"warlock-backend/config"
	"warlock-backend/models"

	"github.com/gin-gonic/gin"
)

type SubjectResult struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

func GetSubjects() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var subjects []models.Subject
		if err := config.DB.Find(&subjects).Error; err != nil {
			ctx.JSON(http.StatusInternalServerError, "An error happened during fetching subjects from DB")
			log.Println("Error during fetching subjects: ", err.Error())
			return
		}

		results := make([]SubjectResult, len(subjects))
		for i, subject := range subjects {
			result := SubjectResult{
				ID:   subject.ID,
				Name: subject.Name,
			}
			results[i] = result
		}

		ctx.JSON(http.StatusOK, results)
	}
}
