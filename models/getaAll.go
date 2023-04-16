package models

import (
	"gopi/db"
)

func GetAllSchools() (s []School, err error) {
	conn, err := db.OpenConnection()
	if err != nil {
		return
	}
	defer conn.Close()

	stmt := `SELECT * FROM schools`
	rows, err := conn.Query(stmt)
	if err != nil {
		return
	}

	for rows.Next() {
		var school School

		err = rows.Scan(&school.ID, &school.Name, &school.Type)
		if err != nil {
			continue
		}

		s = append(s, school)
	}

	return
}
