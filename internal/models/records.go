package models

import (
	"database/sql"
	"time"
)

type Record struct {
	ID         int       `db:"id"`
	Title      string    `db:"title"`
	Artist     string    `db:"artist"`
	CoverImage string    `db:"cover_image"`
	CreatedAt  time.Time `db:"created_at"`
	UpdatedAt  time.Time `db:"updated_at"`
}

type RecordModel struct {
	DB *sql.DB
}
