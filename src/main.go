package main

import (
    "fmt"
    "log"
    "net/http"
    "greet"
)

func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Println(greet.Morning)
    fmt.Fprintf(w, "Hello, my %s!", r.URL.Path[1:])
}

func main() {
    http.HandleFunc("/", handler)
    log.Fatal(http.ListenAndServe(":3000", nil))
}