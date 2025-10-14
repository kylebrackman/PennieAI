package handlers

import (
	"PennieAI/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AnalyzeUnprocessedDocument(c *gin.Context) {
	// Placeholder for analyzing an unprocessed document

	file, err := c.FormFile("document")

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Failed to get document from request",
			"message": err.Error(),
		})
		return
	}

	aiService := services.NewAIService()
	patient, analyzedDocuments, err := services.AnalyzeDocument(file, aiService) // ← Fixed

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Failed to analyze document",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"patient":   patient,
		"documents": analyzedDocuments,
		"count":     len(analyzedDocuments),
		"message":   "Document analyzed successfully",
	})
}
