package main

import (
	"database/sql"
	"html/template"
	"log/slog"
	"net/http"
	"os"

	"github.com/dazakharova/vinyl-condition-tracker/internal/models"
	_ "github.com/mattn/go-sqlite3"
)

type application struct {
	records       models.RecordModel
	recordSides   models.RecordSideModel
	logger        *slog.Logger
	templateCache map[string]*template.Template
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "4001"
	}

	dbPath := os.Getenv("DB_PATH")
	if dbPath == "" {
		dbPath = "./data/pedro.db"
	}

	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	db, err := openDB(dbPath)
	if err != nil {
		logger.Error((err.Error()))
		os.Exit(1)
	}
	defer db.Close()

	templateCache, err := newTemplateCache()
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}

	app := &application{
		records:       models.RecordModel{DB: db},
		recordSides:   models.RecordSideModel{DB: db},
		logger:        logger,
		templateCache: templateCache,
	}

	logger.Info("starting server", slog.String("port", port))

	err = http.ListenAndServe(":"+port, app.routes())
	logger.Error((err.Error()))
	os.Exit(1)
}

func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("sqlite3", dsn)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		db.Close()
		return nil, err
	}

	return db, nil
}
