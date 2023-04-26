package db

import (
	_ "github.com/denisenkom/go-mssqldb"
	"github.com/jmoiron/sqlx"
)

func UseDB() (db *sqlx.DB, err error) {
	db, err = sqlx.Open("sqlserver", "sqlserver://sa:Keng1234@localhost?database=TestDB")
	if err != nil {
		return nil, err
	}
	return db, nil
}
