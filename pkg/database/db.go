package database

import (
	"database/sql"
	_ "github.com/lib/pq"
)

var db *sql.DB

func InitDb() error {
	d, err := sql.Open("postgres", "user=dev password=abcde dbname=passwordManager sslmode=disable")
	if err != nil {
		return err
	}
	db = d
	_, err = db.Exec(schema)
	return err
}
