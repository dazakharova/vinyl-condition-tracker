package main

import "net/http"

func (app *application) routes() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /{$}", app.home)
	mux.HandleFunc("GET /record/view/{id}", recordView)
	mux.HandleFunc("GET /record/create", recordCreate)

	return mux
}
