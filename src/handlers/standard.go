package handlers

import (
	"io"
	"net/http"
)

func Standard(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	io.WriteString(w, `{"maybe": "ok"}`)
}
