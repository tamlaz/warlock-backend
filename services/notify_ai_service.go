package services

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"warlock-backend/models"
)

const documentIngestionUrl string = "/document-ingestion"

func NotifyAiService(document models.Document, userId uint, subjectId uint, topicId uint) {
	go func(doc models.Document) {
		payload := struct {
			UserId       uint   `json:"user_id"`
			SubjectId    uint   `json:"subject_id"`
			TopicId      uint   `json:"topic_id"`
			DocumentId   uint   `json:"document_id"`
			DocumentPath string `json:"document_path"`
			DocumentType string `json:"document_type"`
		}{
			UserId:       userId,
			SubjectId:    subjectId,
			TopicId:      topicId,
			DocumentId:   doc.ID,
			DocumentPath: doc.FilePath,
			DocumentType: doc.DocumentType,
		}

		jsonData, err := json.Marshal(payload)
		if err != nil {
			log.Printf("Failed to marshal json from payload %v", err.Error())
			return
		}

		warlockAiServiceBaseUrl := os.Getenv("WARLOCK_AI_BASE_URL")
		warlockAiApiVersion := os.Getenv("WARLOCK_AI_API_VERSION")
		warlockAiApiPathPrefix := os.Getenv("WARLOCK_AI_API_PATH_PREFIX")
		response, err := http.Post(warlockAiServiceBaseUrl+warlockAiApiVersion+warlockAiApiPathPrefix+documentIngestionUrl,
			"application/json", bytes.NewBuffer(jsonData))
		if err != nil {
			log.Printf("Failed to send request to Warlock AI service: %v", err.Error())
			return
		}
		defer response.Body.Close()

		log.Printf("Async POST request completed with status: %s", response.Status)

	}(document)
}
