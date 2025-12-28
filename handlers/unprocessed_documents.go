package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"PennieAI/config"
	"PennieAI/models"
)

func GetAllUnprocessedDocuments(c *gin.Context) {
	db := config.GetDB()

	var documents []models.UnprocessedDocument
	err := db.Select(&documents, "SELECT * FROM unprocessed_documents ORDER BY created_at DESC")

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Failed to fetch unprocessed documents",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data":    documents,
		"count":   len(documents),
		"message": "Unprocessed documents retrieved successfully",
	})
}
