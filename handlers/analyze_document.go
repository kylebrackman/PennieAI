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

	fileLines, err := services.AnalyzeDocument(file)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Failed to analyze document",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data":    fileLines,
		"count":   len(fileLines),
		"message": "Document analyzed successfully",
	})
}
