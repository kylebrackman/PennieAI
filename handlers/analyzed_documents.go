package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"PennieAI/config"
	"PennieAI/models"
)

// GetAllDocuments retrieves all documents from the database
func GetAllAnalyzedDocuments(c *gin.Context) {
	db := config.GetDB()

	var documents []models.AnalyzedDocument
	err := db.Select(&documents, "SELECT * FROM documents ORDER BY created_at DESC")

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Failed to fetch documents",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data":    documents,
		"count":   len(documents),
		"message": "Documents retrieved successfully",
	})
}

// GetDocumentByID retrieves a single document by ID
func GetDocumentByID(c *gin.Context) {
	db := config.GetDB()

	// Get ID from URL parameter
	idParam := c.Param("id")
	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid document ID format",
		})
		return
	}

	var document models.AnalyzedDocument
	err = db.Get(&document, "SELECT * FROM documents WHERE id = $1", id)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Document not found",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": document,
	})
}

// CreateDocument creates a new document
func CreateDocument(c *gin.Context) {
	db := config.GetDB()

	var req CreateDocumentRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Invalid request format",
			"message": err.Error(),
		})
		return
	}

	var document models.AnalyzedDocument
	query := `
		INSERT INTO documents (title, content, document_type, veterinarian_id, created_at, updated_at) 
		VALUES ($1, $2, $3, $4, NOW(), NOW())
		RETURNING *`

	err := db.Get(&document, query, req.Title, req.Content, req.DocumentType, req.VeterinarianID)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Failed to create document",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"data":    document,
		"message": "Document created successfully",
	})
}

// UpdateDocument updates an existing document
func UpdateDocument(c *gin.Context) {
	db := config.GetDB()

	// Get ID from URL
	idParam := c.Param("id")
	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid document ID format",
		})
		return
	}

	var req UpdateDocumentRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Invalid request format",
			"message": err.Error(),
		})
		return
	}

	var document models.AnalyzedDocument
	query := `
		UPDATE documents 
		SET title = $2, content = $3, document_type = $4, updated_at = NOW()
		WHERE id = $1
		RETURNING *`

	err = db.Get(&document, query, id, req.Title, req.Content, req.DocumentType)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Document not found or failed to update",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data":    document,
		"message": "Document updated successfully",
	})
}

// DeleteDocument deletes a document
func DeleteDocument(c *gin.Context) {
	db := config.GetDB()

	idParam := c.Param("id")
	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid document ID format",
		})
		return
	}

	result, err := db.Exec("DELETE FROM documents WHERE id = $1", id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Failed to delete document",
			"message": err.Error(),
		})
		return
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Document not found",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Document deleted successfully",
	})
}

// Request structs for JSON binding
type CreateDocumentRequest struct {
	Title          string `json:"title" binding:"required"`
	Content        string `json:"content" binding:"required"`
	DocumentType   string `json:"document_type" binding:"required"`
	VeterinarianID int64  `json:"veterinarian_id" binding:"required"`
}

type UpdateDocumentRequest struct {
	Title        string `json:"title" binding:"required"`
	Content      string `json:"content" binding:"required"`
	DocumentType string `json:"document_type" binding:"required"`
}
