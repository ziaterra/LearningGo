package main

import (
	"io"
	"net/http"
)

// Reciever that accepts requests
func index(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "hello string")
}

func main() {
	// Register a function handler
	mux := http.NewServeMux()
	mux.HandleFunc("/", index)
	http.ListenAndServe(":8080", mux) // use default
}
