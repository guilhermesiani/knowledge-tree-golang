package main

import (
	"handlers"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/getall", handlers.GetAll)
	http.HandleFunc("/", handlers.Standard)
	log.Fatal(http.ListenAndServe(":3000", nil))
}
