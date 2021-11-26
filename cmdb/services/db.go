package services

import (
	"database/sql"
)

var db *sql.DB

func InitDb(typ, dsn string) error {
	//db, err := sql.Open(tpy, dsn)
	var err error
	db, err = sql.Open(typ, dsn)
	if err != nil {
		return err
	}
	if err = db.Ping(); err != nil {
		return err
	}
	return nil
}
