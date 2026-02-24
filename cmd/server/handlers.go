package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	records, err := app.records.Latest()
	if err != nil {
		println("DB error!")
		app.serverError(w, r, err)
	}

	data := templateData{
		Records: records,
	}

	files := []string{
		"./ui/html/base.tmpl",
		"./ui/html/pages/home.tmpl",
	}

	ts, err := template.ParseFiles(files...)
	if err != nil {
		app.serverError(w, r, err)
		return
	}

	err = ts.ExecuteTemplate(w, "base", data)
	if err != nil {
		app.serverError(w, r, err)
	}
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
