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

	services.AnalyzeDocument(file)

}
