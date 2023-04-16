package models

import "gopi/db"

func UpdateSchool(id int, s School) (int64, error) {
	conn, err := db.OpenConnection()
	if err != nil {
		return 0, err
	}
	defer conn.Close()

	stmt := `UPDATE schools SET name=$1, type=$2 WHERE id=$3`
	row, err := conn.Exec(stmt, s.Name, s.Type, id)
	if err != nil {
		return 0, err
	}

	return row.RowsAffected()
}
