package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"PennieAI/handlers"
	"PennieAI/middleware"
)

// SetupRoutes configures all application routes
func SetupRoutes(router *gin.Engine) {
	// Basic middlewares
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	router.Use(middleware.CORSMiddleware())

	router.GET("/health", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"status":  "healthy",
			"service": "PennieAI",
			"version": "1.0.0",
			"message": "Veterinary AI assistant is running! 🐾",
		})
	})

	// API versioning
	v1 := router.Group("/api/v1")
	{
		// Document management routes
		documents := v1.Group("/documents")
		{
			documents.GET("", handlers.GetAllDocuments)       // GET /api/v1/documents
			documents.GET("/:id", handlers.GetDocumentByID)   // GET /api/v1/documents/:id
			documents.POST("", handlers.CreateDocument)       // POST /api/v1/documents
			documents.DELETE("/:id", handlers.DeleteDocument) // DELETE /api/v1/documents/:id
		}

		aiTool := v1.Group("/ai_tool")
		{
			aiTool.GET("/test", handlers.TestAiService)
			aiTool.GET("/model_version", handlers.GetAiModelVersion)
		}

		// Unprocessed document routes
		unprocessedDocuments := v1.Group("/unprocessed")
		{
			unprocessedDocuments.GET("", handlers.GetAllUnprocessedDocuments)          // GET /api/v1/unprocessed
			unprocessedDocuments.POST("/analyze", handlers.AnalyzeUnprocessedDocument) // POST /api/v1/unprocessed/analyze
		}

	}

	// Catch-all route for API documentation or 404
	router.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{
			"error":   "Endpoint not found",
			"message": "Visit /health for status or /api/v1 for API routes",
		})
	})
}
