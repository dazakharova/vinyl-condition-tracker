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

func (m *RecordSideModel) Get(recordID int) ([]RecordSide, error) {
	stmt := "SELECT id, record_id, name FROM record_sides WHERE record_id = ?"
	rows, err := m.DB.Query(stmt, recordID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var recordSides []RecordSide

	for rows.Next() {
		var rs RecordSide
		err := rows.Scan(
			&rs.ID,
			&rs.RecordID,
			&rs.Name,
		)
		if err != nil {
			return nil, err
		}

		recordSides = append(recordSides, rs)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return recordSides, nil
}
