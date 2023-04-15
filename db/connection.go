package db

import (
	"database/sql"
	"fmt"
	"gopi/config"

	_ "github.com/lib/pq"
)

func OpenConnection() (*sql.DB, error) {
	conf := config.GetDBConfig()

	sc := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		conf.Host, conf.Port, conf.User, conf.Password, conf.Database,
	)

	db, err := sql.Open("postgres", sc)
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	return db, err
}
