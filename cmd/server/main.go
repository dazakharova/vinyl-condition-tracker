package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello from Vinyl Record Condition Tracker"))
}

func recordView(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}

	msg := fmt.Sprintf("Display a specific record with ID %d", id)
	w.Write([]byte(msg))
}

func recordCreate(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Display a form for creating a new record"))
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/{$}", home)
	mux.HandleFunc("/record/view/{id}", recordView)
	mux.HandleFunc("/record/create", recordCreate)

	log.Println("Starting server on port :4000")

	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
