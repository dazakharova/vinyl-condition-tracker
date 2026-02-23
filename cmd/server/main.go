package main

import (
	"log"
	"net/http"
	"os"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "4001"
	}

	mux := http.NewServeMux()
	mux.HandleFunc("GET /{$}", home)
	mux.HandleFunc("GET /record/view/{id}", recordView)
	mux.HandleFunc("GET /record/create", recordCreate)

	log.Printf("Starting server on port :%s", port)

	err := http.ListenAndServe(":"+port, mux)
	log.Fatal(err)
}
