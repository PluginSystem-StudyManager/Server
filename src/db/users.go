package db

import (
	"database/sql"
	"errors"
	"log"
)

func AddUser(username string, password string) (sql.Result, error) {
	return insert("INSERT INTO users(username, password) values (?, ?)", username, password)
}

func UpdateToken(username string, token string, ttl string) (sql.Result, error) {
	return insert(`UPDATE users SET token=?, token_ttl=? WHERE username=?`, token, ttl, username)
}

func UserIdByToken(token string) (int, error) {
	rows, err := db.Query(`SELECT id FROM users WHERE token=? AND token_ttl>=datetime('now')`, token)
	if err != nil {
		log.Fatal(err)
		return -1, err
	}
	defer rows.Close()
	for rows.Next() {
		var id int
		err = rows.Scan(&id)
		if err != nil {
			log.Fatal(err)
			return -1, err
		}
		return id, nil
	}
	return -1, errors.New("no valid token found: " + token)
}
