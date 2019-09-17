package main

import (
    "log"
    "net/http"
    "handlers"
)

func main() {
    http.HandleFunc("/", handlers.Standard)
    log.Fatal(http.ListenAndServe(":3000", nil))
}