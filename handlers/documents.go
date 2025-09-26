package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"PennieAI/config"
	"PennieAI/models"
)

// GetAllDocuments retrieves all documents from the database
func GetAllDocuments(context *gin.Context) {
	db := config.GetDB()

	var documents []models.Document
	err := db.Select(&documents, "SELECT * FROM documents ORDER BY created_at DESC")

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Failed to fetch documents",
			"message": err.Error(),
		})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"data":    documents,
		"count":   len(documents),
		"message": "Documents retrieved successfully",
	})
}

// GetDocumentByID retrieves a single document by ID
func GetDocumentByID(context *gin.Context) {
	db := config.GetDB()

	// Get ID from URL parameter
	idParam := context.Param("id")
	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid document ID format",
		})
		return
	}

	var document models.Document
	err = db.Get(&document, "SELECT * FROM documents WHERE id = $1", id)

	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{
			"error": "Document not found",
		})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"data": document,
	})
}

// CreateDocument creates a new document
func CreateDocument(context *gin.Context) {
	db := config.GetDB()

	var req CreateDocumentRequest
	if err := context.ShouldBindJSON(&req); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"error":   "Invalid request format",
			"message": err.Error(),
		})
		return
	}

	var document models.Document
	query := `
		INSERT INTO documents (title, content, document_type, veterinarian_id, created_at, updated_at) 
		VALUES ($1, $2, $3, $4, NOW(), NOW())
		RETURNING *`

	err := db.Get(&document, query, req.Title, req.Content, req.DocumentType, req.VeterinarianID)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Failed to create document",
			"message": err.Error(),
		})
		return
	}

	context.JSON(http.StatusCreated, gin.H{
		"data":    document,
		"message": "Document created successfully",
	})
}

// UpdateDocument updates an existing document
func UpdateDocument(context *gin.Context) {
	db := config.GetDB()

	// Get ID from URL
	idParam := context.Param("id")
	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid document ID format",
		})
		return
	}

	var req UpdateDocumentRequest
	if err := context.ShouldBindJSON(&req); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"error":   "Invalid request format",
			"message": err.Error(),
		})
		return
	}

	var document models.Document
	query := `
		UPDATE documents 
		SET title = $2, content = $3, document_type = $4, updated_at = NOW()
		WHERE id = $1
		RETURNING *`

	err = db.Get(&document, query, id, req.Title, req.Content, req.DocumentType)

	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{
			"error": "Document not found or failed to update",
		})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"data":    document,
		"message": "Document updated successfully",
	})
}

// DeleteDocument deletes a document
func DeleteDocument(context *gin.Context) {
	db := config.GetDB()

	idParam := context.Param("id")
	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid document ID format",
		})
		return
	}

	result, err := db.Exec("DELETE FROM documents WHERE id = $1", id)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Failed to delete document",
			"message": err.Error(),
		})
		return
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		context.JSON(http.StatusNotFound, gin.H{
			"error": "Document not found",
		})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"message": "Document deleted successfully",
	})
}

func AnalyzeDocument(context *gin.Context) {
	// Placeholder for future AI analysis implementation
	context.JSON(http.StatusNotImplemented, gin.H{
		"message": "Document analysis feature is not yet implemented",
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
