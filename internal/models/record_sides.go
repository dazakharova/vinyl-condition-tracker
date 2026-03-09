package models

import (
	"database/sql"
)

type RecordSide struct {
	ID       int    `db:"id"`
	RecordID int    `db:"record_id"`
	Name     string `db:"name"`
}

type RecordSideModel struct {
	DB *sql.DB
}

func (m *RecordSideModel) Insert(recordID int, name string) (int, error) {
	stmt := "INSERT INTO record_sides (record_id, name) VALUES (?, ?)"
	result, err := m.DB.Exec(stmt, recordID, name)
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), nil
}
