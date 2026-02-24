package main

import (
	"time"

	"github.com/dazakharova/vinyl-condition-tracker/internal/models"
)

type templateData struct {
	Records []models.Record
}

func humanDate(t time.Time) string {
	if t.IsZero() {
		return ""
	}

	return t.UTC().Format("02 Jan 2006 at 15:04")
}
