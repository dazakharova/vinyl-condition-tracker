package main

import (
	"fmt"
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
