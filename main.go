package main

import (
	"fmt"
	"log"
	"os"

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
	product := Product{"pc", 699.00, true}

	id := insertProduct(db, product)
	fmt.Printf("Product inserted with id: %d\n", id)

	// query Single row
	// var p Product

	// query := "SELECT name, price, available FROM products WHERE id = $1"
	// err = db.QueryRow(query, id).Scan(&p.Name, &p.Price, &p.Available)

	// if err != nil {
	// 	if err == sql.ErrNoRows {
	// 		log.Fatalf("No product found with the id %d\n", id)
	// 	} else {
	// 		log.Fatalf("Error scanning product: %q", err)
	// 	}
	// }

	// fmt.Printf("Name: %s\n", p.Name)
	// fmt.Printf("Price: %.2f\n", p.Price)       // %.2f - prints float with 2 decimal places
	// fmt.Printf("Available: %t\n", p.Available) // %t - prints true or false

	// query Multiple rows
	data := []Product{}
	rows, err := db.Query("SELECT name, price, available FROM products")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var p Product
		if err := rows.Scan(&p.Name, &p.Price, &p.Available); err != nil {
			log.Fatal(err)
		}
		data = append(data, p)
	}
	fmt.Println(data)
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
