package models

import (
	"database/sql"
	"errors"
	"time"
)

var ErrNoRecord = errors.New("models: no matching record found")

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

func (m *RecordModel) Latest() ([]Record, error) {
	stmt := "SELECT id, title, artist, created_at, updated_at FROM records ORDER BY created_at DESC LIMIT 10"
	rows, err := m.DB.Query(stmt)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var records []Record

	for rows.Next() {
		var r Record
		err := rows.Scan(
			&r.ID,
			&r.Title,
			&r.Artist,
			&r.CreatedAt,
			&r.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}

		records = append(records, r)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return records, nil
}

func (m *RecordModel) Insert(title, artist string) (int, error) {
	stmt := "INSERT INTO records (title, artist) VALUES (?, ?)"
	result, err := m.DB.Exec(stmt, title, artist)
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), nil
}

func (m *RecordModel) Get(id int) (Record, error) {
	stmt := "SELECT id, title, artist, created_at, updated_at FROM records WHERE id = ?"
	row := m.DB.QueryRow(stmt, id)

	var r Record
	err := row.Scan(&r.ID, &r.Title, &r.Artist, &r.CreatedAt, &r.UpdatedAt)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return Record{}, ErrNoRecord
		}
		return Record{}, err
	}

	return r, nil
}
