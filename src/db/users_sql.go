//+build

package db

import (
	"errors"
	"log"
)

func AddUser(username string, password string, firstName string, lastName string, email string) error {
	_, err := insert("INSERT INTO users(username, password, firstName, lastName, e_mail) values (?, ?, ?, ?, ?)", username, password, firstName, lastName, email)
	return err
}

func CheckCredentials(username string, password string) (bool, error) {

	rows, err := db.Query(`SELECT username, password FROM users where  username=? AND password=?`, username, password)
	if err != nil {
		log.Printf("Query Error: %v", err)
		return false, err
	}

	defer rows.Close()
	for rows.Next() {

		var uName string
		var pw string

		err := rows.Scan(&uName, &pw)
		if err != nil {
			log.Printf("Query Error: %v", err)
			return false, err
		}
		if username == uName && password == pw {
			return true, nil
		}
	}
	return false, nil
}

func UsernameAvailable(username string) (bool, error) {

	rows, err := db.Query(`SELECT username FROM users where  username=?`, username)

	if err != nil {
		log.Printf("Query Error: %v", err)
		return false, err
	}

	defer rows.Close()
	for rows.Next() {

		var uName string

		err := rows.Scan(&uName)
		if err != nil {
			log.Printf("Query Error: %v", err)
			return false, err
		}
		if username == uName {
			return false, nil
		}
	}
	return true, nil
}

func UpdateToken(username string, token string, ttl string) error {
	_, err := insert(`UPDATE users SET token=?, token_ttl=? WHERE username=?`, token, ttl, username)
	return err
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

func UserIdByUsername(username string) (int, error) {
	rows, err := db.Query(`SELECT id FROM users WHERE username=?`, username)
	if err != nil {
		log.Fatal(err)
		return -1, err
	}
	defer rows.Close()
	for rows.Next() {
		var id int
		err = rows.Scan(&id)
		if err != nil {
			log.Print(err)
			return -1, err
		}
		return id, nil
	}
	return -1, errors.New("No user with username: " + username)
}
