package handler

import (
    "fmt",
    "net/http"
)

func standard(w http.ResponseWriter, r *http.Request) {
    fmt.Println(greet.Morning)
    fmt.Fprintf(w, "Hello, my %s!", r.URL.Path[1:])
}