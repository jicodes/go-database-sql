package main

import (
	"log"
	"os"

	"database/sql"

	_ "github.com/lib/pq"

	"github.com/joho/godotenv"
)

func main() {
	// Load the .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	// Create a new connection to the database
	connStr := os.Getenv("CONNECTION_STRING")
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("Error opening database: %q", err)
	}
	defer db.Close()

	// Check if the connection is alive
	if err = db.Ping(); err != nil {
		log.Fatalf("Error pinging database: %q", err)
	}

  // Create the products table
  createProductsTable(db)
}

func createProductsTable(db *sql.DB) {
	// Create a new table
	query := `CREATE TABLE IF NOT EXISTS products (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    price NUMERIC(6, 2) NOT NULL,
    available BOOLEAN NOT NULL DEFAULT TRUE,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
  );`

	_, err := db.Exec(query)
	if err != nil {
		log.Fatalf("Error creating table: %q", err)
	}
}
