package main

import (
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /{$}", home)
	mux.HandleFunc("GET /record/view/{id}", recordView)
	mux.HandleFunc("GET /record/create", recordCreate)

	log.Println("Starting server on port :4000")

	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
