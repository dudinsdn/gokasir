package main

import (
	"log"
	"net/http"

	"github.com/dudinsdn/gokasir/internal/config"
	"github.com/dudinsdn/gokasir/internal/handler"
	"github.com/dudinsdn/gokasir/internal/repository"
	"github.com/dudinsdn/gokasir/internal/usecase"
)

func main() {
	db := config.NewPostgresDB()
	productRepo := repository.NewProductRepository(db)
	productUsecase := usecase.NewProductUsecase(productRepo)
	productHandler := handler.NewProductHandler(productUsecase)

	mux := http.NewServeMux()
	mux.HandleFunc("/api/products", productHandler.ListProducts)

	log.Println("starting server on :8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
