package config

import (
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
)

func RunMigrations() error {

	environment := os.Getenv("ENV")
	if environment == "" {
		environment = "development"
	}
	if environment != "development" {
		log.Println("Environment is not development, skipping automatic migrations.")
		return nil
	}
	db := GetDB()

	/*
	  repository.DB is accessing the underlying database/sql connection from sqlx
	  postgres.WithInstance wraps this connection in migration-specific logic
	*/
	driver, err := postgres.WithInstance(db.DB, &postgres.Config{})
	if err != nil {
		return fmt.Errorf("failed to create migration driver: %w", err)
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file://migrations",
		"postgres", // Database driver name
		driver,
	)
	if err != nil {
		return fmt.Errorf("failed to create migration instance: %w", err)
	}

	log.Println("Starting database migrations...")
	err = m.Up()
	if err != nil && !errors.Is(err, migrate.ErrNoChange) {
		return fmt.Errorf("migration failed: %w", err)
	}

	if errors.Is(err, migrate.ErrNoChange) {
		log.Println("✅ Database is up to date (no migrations to run)")
	} else {
		log.Println("✅ Migrations completed successfully")
	}

	return nil
}
