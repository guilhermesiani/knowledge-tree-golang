package handlers

import (
    "net/http"
    "io"
)

func Standard(w http.ResponseWriter, r *http.Request) {
    w.WriteHeader(http.StatusOK)
    w.Header().Set("Content-Type", "application/json")

    io.WriteString(w, `{"maybe": "ok"}`)
}