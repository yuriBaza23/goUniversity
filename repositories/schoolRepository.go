package repositories

import "gopi/models"

func DeleteSchool(id int) (int64, error) {
	conn, err := OpenConnection()
	if err != nil {
		return 0, err
	}
	defer conn.Close()

	stmt := `DELETE FROM schools WHERE id=$1`
	row, err := conn.Exec(stmt, id)
	if err != nil {
		return 0, err
	}

	return row.RowsAffected()
}

func InsertSchool(school models.School) (id int, err error) {
	err = school.VerifyType()
	if err != nil {
		return
	}

	conn, err := OpenConnection()
	if err != nil {
		return
	}
	defer conn.Close()

	stmt := `INSERT INTO schools (name, type) VALUES ($1, $2) RETURNING id`

	err = conn.QueryRow(stmt, school.Name, school.Type).Scan(&id)

	return
}

func GetSchool(id int) (s models.School, err error) {
	conn, err := OpenConnection()
	if conn != nil {
		return
	}
	defer conn.Close()

	stmt := `SELECT * FROM schools WHERE id=$1`

	row := conn.QueryRow(stmt, id)

	err = row.Scan(&s.ID, &s.Name, &s.Type)

	return
}

func GetAllSchools() (s []models.School, err error) {
	conn, err := OpenConnection()
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
		var school models.School

		err = rows.Scan(&school.ID, &school.Name, &school.Type)
		if err != nil {
			continue
		}

		s = append(s, school)
	}

	return
}

func UpdateSchool(id int, s models.School) (int64, error) {
	conn, err := OpenConnection()
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
