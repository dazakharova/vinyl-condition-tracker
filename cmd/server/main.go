package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/dazakharova/vinyl-condition-tracker/internal/models"
	_ "github.com/mattn/go-sqlite3"
)

type application struct {
	records models.RecordModel
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

	db, err := openDB(dbPath)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	app := &application{
		records: models.RecordModel{DB: db},
	}

	log.Printf("Starting server on port :%s", port)

	err = http.ListenAndServe(":"+port, app.routes())
	log.Fatal(err)
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
