package main

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/lib/pq"
)

func main() {
	dsn := os.Getenv("DATABASE_URL")
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	migration, err := os.ReadFile("migrations/001_create_products.sql")
	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec(string(migration))
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Migration executed successfully")
}
