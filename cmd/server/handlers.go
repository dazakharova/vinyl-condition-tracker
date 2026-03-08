package main

import (
	"errors"
	"fmt"
	"html/template"
	"net/http"
	"strconv"

	"github.com/dazakharova/vinyl-condition-tracker/internal/models"
	"github.com/dazakharova/vinyl-condition-tracker/internal/validator"
)

type RecordCreateForm struct {
	Title               string `json:"title"`
	Artist              string `json:"artist"`
	validator.Validator `form:"-"`
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

func (app *application) recordView(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}

	record, err := app.records.Get(id)
	if err != nil {
		if errors.Is(err, models.ErrNoRecord) {
			http.NotFound(w, r)
		} else {
			app.serverError(w, r, err)
		}
		return
	}

	data := app.newTemplateData()
	data.Record = record

	app.render(w, r, http.StatusOK, "view.tmpl", data)
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

func (app *application) recordCreatePost(w http.ResponseWriter, r *http.Request) {
	var form RecordCreateForm

	err := r.ParseForm()
	if err != nil {
		app.clientError(w, http.StatusBadRequest)
		return
	}

	form.Title = r.PostForm.Get("title")
	form.Artist = r.PostForm.Get("artist")

	form.CheckField(form.NotBlank(form.Title), "title", "The field can not be blank")
	form.CheckField(form.NotBlank(form.Artist), "artist", "The field can not be blank")

	if !form.Valid() {
		var data templateData
		data.Form = form
		app.render(w, r, http.StatusUnprocessableEntity, "create.tmpl", data)
		return
	}

	_, err = app.records.Insert(form.Title, form.Artist)
	if err != nil {
		app.serverError(w, r, err)
		return
	}

	http.Redirect(w, r, fmt.Sprint("/"), http.StatusSeeOther)
	//http.Redirect(w, r, fmt.Sprintf("/record/view/%d", id), http.StatusSeeOther)
}
