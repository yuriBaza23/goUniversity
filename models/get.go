package models

import (
	"gopi/db"
)

func GetSchool(id int) (s School, err error) {
	conn, err := db.OpenConnection()
	if conn != nil {
		return
	}
	defer conn.Close()

	stmt := `SELECT * FROM schools WHERE id=$1`

	row := conn.QueryRow(stmt, id)

	err = row.Scan(&s.ID, &s.Name, &s.Type)

	return
}
