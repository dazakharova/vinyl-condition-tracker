package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
)

//type RecordCreateForm struct {
//	Title               string `json:"title"`
//	Artist              string `json:"artist"`
//	validator.Validator `form:"-"`
//}

var functions = template.FuncMap{
	"humanDate": humanDate,
}

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

	ts, err := template.New("home").Funcs(functions).ParseFiles(files...)
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

func (app *application) recordCreate(w http.ResponseWriter, r *http.Request) {
	files := []string{
		"./ui/html/base.tmpl",
		"./ui/html/pages/create.tmpl",
	}
	ts, err := template.ParseFiles(files...)
	if err != nil {
		app.serverError(w, r, err)
		return
	}

	err = ts.ExecuteTemplate(w, "base", nil)
	if err != nil {
		app.serverError(w, r, err)
	}
}
