package main

import (
	"database/sql"
	"log/slog"
	"net/http"
	"os"

	"github.com/dazakharova/vinyl-condition-tracker/internal/models"
	_ "github.com/mattn/go-sqlite3"
)

type application struct {
	records models.RecordModel
	logger  *slog.Logger
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

	app := &application{
		records: models.RecordModel{DB: db},
		logger:  logger,
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
