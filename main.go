package main

import (
	"database/sql"
  _ "github.com/lib/pq"
	"log"
	"os"

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
}
