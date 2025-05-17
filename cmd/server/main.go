package main

import (
	"log"

	"github.com/dudinsdn/gokasir/internal/config"
)

func main() {
	config.NewPostgresDB()
	log.Println("Connected to the database successfully")
}
