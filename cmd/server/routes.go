package main

import "net/http"

func (app *application) routes() *http.ServeMux {
	mux := http.NewServeMux()

	fileServer := http.FileServer(http.Dir("./ui/static"))
	mux.Handle("GET /static/", http.StripPrefix("/static/", fileServer))

	mux.HandleFunc("GET /{$}", app.home)
	mux.HandleFunc("GET /record/view/{id}", app.recordView)
	mux.HandleFunc("GET /record/create", app.recordCreate)
	mux.HandleFunc("POST /record/create", app.recordCreatePost)

	return mux
}
