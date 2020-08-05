package db

import (
	"database/sql"
	"errors"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

var db *sql.DB

func Open() {
	dbLoc, err := sql.Open("sqlite3", "../dist/foo.db")
	if err != nil {
		log.Fatal(err)
		return
	}
	db = dbLoc
}

func Close() {
	_ = db.Close()
}

// Users

func addUser(username string, password string) error {
	return insert("INSERT INTO users(username, password) values (?, ?)", username, password)
}

func updateToken(username string, token string, ttl string) error {
	return insert(`UPDATE users SET token=?, token_ttl=? WHERE username=?`, token, ttl, username)
}

func userIdByToken(token string) (int, error) {
	rows, err := db.Query(`SELECT id FROM users WHERE token=? AND token_ttl>=datetime('now')`, token)
	if err != nil {
		log.Fatal(err)
		return -1, err
	}
	for rows.Next() {
		var id int
		err = rows.Scan(id)
		if err != nil {
			log.Fatal(err)
			return -1, err
		}
		return id, nil
	}
	return -1, errors.New("no valid token found: " + token)
}

func insert(statement string, args ...interface{}) error {
	conn, err := db.Begin()
	if err != nil {
		log.Fatal(err)
		return err
	}
	stmt, err := db.Prepare(statement)
	if err != nil {
		log.Fatal(err)
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(args)
	if err != nil {
		log.Fatal(err)
		return err
	}
	_ = conn.Commit()
	return nil
}
