package db

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

func insert(statement string, args ...interface{}) (sql.Result, error) {
	stmt, err := db.Prepare(statement)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	res, err := stmt.Exec(args...)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	err = stmt.Close()
	if err != nil {
		log.Printf("error closing stmt: %v\n", err)
	}
	return res, nil
}
