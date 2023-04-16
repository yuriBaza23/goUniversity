package models

import "gopi/db"

func InsertSchool(school School) (id int, err error) {
	err = school.VerifyType()
	if err != nil {
		return
	}

	conn, err := db.OpenConnection()
	if err != nil {
		return
	}
	defer conn.Close()

	stmt := `INSERT INTO schools (name, type) VALUES ($1, $2) RETURNING id`

	err = conn.QueryRow(stmt, school.Name, school.Type).Scan(&id)

	return
}
