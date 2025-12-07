package config

import (
	"log"
	"os"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq" // PostgreSQL driver
)

var DB *sqlx.DB

func InitDatabase() {
	var err error

	databaseUrl := os.Getenv("DATABASE_URL")
	if databaseUrl == "" {
		log.Fatal("DATABASE_URL environment variable is required")
	}

	DB, err = sqlx.Connect("postgres", databaseUrl)

	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// Configure connection pool (optional but recommended)
	DB.SetMaxIdleConns(10)
	DB.SetMaxOpenConns(100)

	if err := DB.Ping(); err != nil {
		log.Fatal("Failed to ping database:", err)
	}

	log.Println("✅ Database connected successfully")
}

func GetDB() *sqlx.DB {
	if DB == nil {
		log.Fatal("Database not initialized. Call InitDatabase() first.")
	}
	return DB
}

func CloseDatabase() error {
	if DB != nil {
		return DB.Close()
	}
	return nil
}
