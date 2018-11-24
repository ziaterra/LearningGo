package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

// Defined handler
func index(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Missy ho!"))
}

func main() {
	// initialize router
	r := chi.NewRouter()

	// Set up Middleware
	r.Use(middleware.Logger)

	// Set up router handlers
	r.Get("/", index)

	// Listen on router mux
	err := http.ListenAndServe(":8081", r)

	// check for any erroes serving the http port
	if err != nil {
		fmt.Println(err)
	}
}
