package main

import (
	"handlers"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/health/ping", handlers.Health)
	http.HandleFunc("/products", handlers.GetAllProducts)
	// http.HandleFunc("/products", handlers.AddProduct)
	log.Fatal(http.ListenAndServe(":3000", nil))
}
