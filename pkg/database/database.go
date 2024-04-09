package database

import (
	"fmt"
	"log"
	"os"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq" // PostgreSQL driver
)

// DB is the database connection handle
var DB *sqlx.DB

// Initialize connects to the database using environment variables and stores the connection handle in DB.
func Initialize() (*sqlx.DB, error) {
	// Construct the DSN (Data Source Name) from environment variables
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"))

	// Connect to the database using sqlx
	db, err := sqlx.Connect("postgres", dsn)
	if err != nil {
		return nil, fmt.Errorf("error connecting to the database: %w", err)
	}

	// Optional: verify the connection with a ping
	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("error verifying connection to the database: %w", err)
	}

	log.Println("Successfully connected to the database")
	return db, nil
}
