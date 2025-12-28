package handlers

import (
	"PennieAI/models"
	"PennieAI/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AnalyzeResponse struct {
	Message   string                    `json:"message"`
	Count     int                       `json:"count"`
	Patient   *models.Patient           `json:"patient"`
	Documents []models.AnalyzedDocument `json:"documents"`
}

func AnalyzeUnprocessedDocument(c *gin.Context) {
	file, err := c.FormFile("document")

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Failed to get document from request",
			"message": err.Error(),
		})
		return
	}

	aiService := services.NewAIService()
	patient, analyzedDocuments, err := services.AnalyzeDocument(file, aiService)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Failed to analyze document",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, AnalyzeResponse{
		Message:   "Document analyzed successfully",
		Count:     len(analyzedDocuments),
		Patient:   patient,
		Documents: analyzedDocuments,
	})
}
