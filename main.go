package main

import (
	"log"
	"os"
  "fmt"

	"database/sql"

	_ "github.com/lib/pq"

	"github.com/joho/godotenv"
)

type Product struct {
	Name      string
	Price     float64
	Available bool
}

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
	// Insert a new product
	product := Product{"egg", 2.99, true}

	id := insertProduct(db, product)
	fmt.Printf("Product inserted with id: %d\n", id)
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

func insertProduct(db *sql.DB, product Product) int {
	// Insert a new product
	query := `INSERT INTO products (name, price, available) VALUES ($1, $2, $3) RETURNING id;`
	var id int

	err := db.QueryRow(query, product.Name, product.Price, product.Available).Scan(&id)
	if err != nil {
		log.Fatalf("Error inserting product: %q", err)
	}
	return id
}