package main

import (
	"html/template"
	"path/filepath"
	"time"

	"github.com/dazakharova/vinyl-condition-tracker/internal/models"
)

type templateData struct {
	Records     []models.Record
	Record      models.Record
	RecordSides int
	Form        any
}

var functions = template.FuncMap{
	"humanDate": humanDate,
}

func humanDate(t time.Time) string {
	if t.IsZero() {
		return ""
	}

	return t.UTC().Format("02 Jan 2006 at 15:04")
}

func newTemplateCache() (map[string]*template.Template, error) {
	cache := map[string]*template.Template{}
	pages, err := filepath.Glob("ui/html/pages/*.tmpl")
	if err != nil {
		return nil, err
	}

	for _, page := range pages {
		name := filepath.Base(page)
		patterns := []string{
			"ui/html/base.tmpl",
			page,
		}

		ts, err := template.New(name).Funcs(functions).ParseFiles(patterns...)
		if err != nil {
			return nil, err
		}

		cache[name] = ts
	}

	return cache, nil
}
