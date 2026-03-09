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
	Sides               string
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
	data := app.newTemplateData()
	data.Form = RecordCreateForm{
		Sides: "2",
	}

	files := []string{
		"./ui/html/base.tmpl",
		"./ui/html/pages/create.tmpl",
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

func (app *application) recordCreatePost(w http.ResponseWriter, r *http.Request) {
	var form RecordCreateForm

	err := r.ParseForm()
	if err != nil {
		app.clientError(w, http.StatusBadRequest)
		return
	}

	form.Title = r.PostForm.Get("title")
	form.Artist = r.PostForm.Get("artist")
	form.Sides = r.PostForm.Get("sides")

	form.CheckField(form.NotBlank(form.Title), "title", "The field can not be blank")
	form.CheckField(form.NotBlank(form.Artist), "artist", "The field can not be blank")

	form.CheckField(form.NotBlank(form.Sides), "sides", "Sides amount cannot be blank")
	form.CheckField(form.IsInt(form.Sides), "sides", "Sides amount must be a whole number")
	form.CheckField(form.GreaterThan(form.Sides, 0), "sides", "Sides amount must be greater than 0")
	form.CheckField(form.IsEven(form.Sides), "sides", "Sides amount must be an even number")

	if !form.Valid() {
		data := app.newTemplateData()
		data.Form = form
		app.render(w, r, http.StatusUnprocessableEntity, "create.tmpl", data)
		return
	}

	sidesCount, err := strconv.Atoi(form.Sides)
	if err != nil {
		app.serverError(w, r, err)
		return
	}

	sideNames := generateSideNames(sidesCount)

	recordID, err := app.records.Insert(form.Title, form.Artist)
	if err != nil {
		app.serverError(w, r, err)
		return
	}

	for _, name := range sideNames {
		_, err = app.recordSides.Insert(strconv.Itoa(recordID), name)
		if err != nil {
			app.serverError(w, r, err)
			return
		}
	}
	//
	//http.Redirect(w, r, fmt.Sprint("/"), http.StatusSeeOther)
	http.Redirect(w, r, fmt.Sprintf("/record/view/%d", recordID), http.StatusSeeOther)
}
