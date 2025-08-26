package config

import (
	"log"
	"os"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq" // PostgreSQL driver
)

/*
*
 - sqlx.DB is the actual database struct
 - *sqlx.DB is a pointer to the database struct
 - var DB *sqlx.DB makes DB a package-level variable that can be accessed from other files in the same package
 - We use a pointer because database connections are expensive resources, and we want to share the same connection pool across the application
 - Pointers let us pass around references, not copies, which is more efficient.
*/

// Package-level variable
var DB *sqlx.DB

// InitDatabase initializes the database connection
// Initialize function
func InitDatabase() {
	var err error

	// Get database URL from environment
	databaseUrl := os.Getenv("DATABASE_URL")
	if databaseUrl == "" {
		log.Fatal("DATABASE_URL environment variable is required")
	}

	/**
	sqlx.Connect returns TWO values: (*sqlx.DB, error)
	We assign them to TWO variables at once
	Connect to the db
	*/
	DB, err = sqlx.Connect("postgres", databaseUrl)
	/**
	This is equivalent to:
		result := sqlx.Connect("postgres", dsn)
		DB = result.db    // First return value
		err = result.err  // Second return value

		Go functions often return (value, error):
			Pattern: value, error := someFunction()
	*/
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// Configure connection pool (optional but recommended)
	DB.SetMaxIdleConns(10)
	DB.SetMaxOpenConns(100)

	// Test the connection
	if err := DB.Ping(); err != nil {
		log.Fatal("Failed to ping database:", err)
	}

	log.Println("✅ Database connected successfully")
}

// GetDB returns the database instance
// Why not use DB directly? Encapsulation and future flexibility. Safety check (prevents nil pointer crashes), Single place to handle "DB not initialized" errors, better error messages
// Getter function
func GetDB() *sqlx.DB {
	if DB == nil {
		log.Fatal("Database not initialized. Call InitDatabase() first.")
	}
	return DB
}

// CloseDatabase closes the database connection
// Cleanup function
func CloseDatabase() error {
	if DB != nil {
		return DB.Close()
	}
	return nil
}
