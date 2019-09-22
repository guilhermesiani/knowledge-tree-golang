package main

import (
	"app/handlers"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/health/ping", handlers.Health).Methods("GET")
	router.HandleFunc("/products", handlers.GetAllProducts).Methods("GET")
	// http.HandleFunc("/products", handlers.AddProduct)
	log.Fatal(http.ListenAndServe(":3000", router))
}
