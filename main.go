package main

import (
	"PennieAI/models"
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"log"
	"os"
)

func main() {
	// Load environment variables
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Get database URL from environment
	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		log.Fatal("DATABASE_URL environment variable is required")
	}

	// Connect
	db, err := sqlx.Connect("postgres", dsn)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Test connection
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	// Query documents
	var documents []models.Document
	err = db.Select(&documents, "SELECT * FROM documents")
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Retrieved %d documents\n", len(documents))
}
