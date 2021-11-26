package services

import (
	"database/sql"
)

var db *sql.DB

func InitDb(tpy, dsn string) error {
	db, err := sql.Open(tpy, dsn)
	if err != nil {
		return err
	}
	if err = db.Ping(); err != nil {
		return err
	}
	return nil
}
