package main

import (
    "fmt"
    "log"
    "net/http"
    "greet"
)

func main() {
    http.HandleFunc("/", handler)
    log.Fatal(http.ListenAndServe(":3000", nil))
}