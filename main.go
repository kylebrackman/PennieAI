package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"PennieAI/config"
	"PennieAI/routes"
)

func main() {
	// Load environment variables
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found, using system environment variables")
	}

	// Initialize database connection
	config.InitDatabase()

	// Gracefully close database connection when main exits
	defer func() {
		if err := config.CloseDatabase(); err != nil {
			log.Printf("Error closing database: %v", err)
		}
	}()

	// Set Gin mode based on environment
	if os.Getenv("GIN_MODE") == "release" {
		gin.SetMode(gin.ReleaseMode)
	}

	// Create Gin router
	router := gin.Default()

	// Setup all routes
	routes.SetupRoutes(router)

	// Get port from environment or default
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	
	// Start server
	log.Printf("🐾 PennieAI API starting on port %s...", port)
	log.Printf("📍 Health check: http://localhost:%s/health", port)
	log.Printf("📍 API docs: http://localhost:%s/api/v1", port)

	if err := router.Run(":" + port); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
