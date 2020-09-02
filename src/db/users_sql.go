package db

import (
	"errors"
	"log"
	"server/utils"
)

func AddUser(username string, password string, firstName string, lastName string, email string) error {
	permanentToken := utils.CreateToken()
	_, err := insert(
		"INSERT INTO users(username, password, firstName, lastName, e_mail, permanent_token) values (?, ?, ?, ?, ?, ?)",
		username, password, firstName, lastName, email, permanentToken)
	return err
}

func AddDebugUser() {
	_, _ = insert(
		"INSERT INTO users(username, password, firstName, lastName, e_mail, permanent_token, token, token_ttl) values (?, ?, ?, ?, ?, ?, ?, ?)",
		"John", "12345", "John", "Ross", "john@ross.com", "12345", "12345", "2099")
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
	var id int
	err := xByY("id", "token=? AND token_ttl>=datetime('now')", token, &id)
	return id, err
}

func UserIdByUsername(username string) (int, error) {
	var id int
	err := xByY("id", "username=?", username, &id)
	return id, err
}

type User struct {
	Username string
}

func UserByToken(token string) (User, error) {
	var user User
	err := xByY("username", "token=? AND token_ttl>=datetime('now')", token, &user.Username)
	return user, err
}

func UserIdByPermanentToken(token string) (int, error) {
	var id int
	err := xByY("id", "permanent_token=?", token, &id)
	return id, err
}

func PermanentTokenByUsername(username string) (string, error) {
	var token string
	err := xByY("permanent_token", "username=?", username, &token)
	return token, err
}

func xByY(selectString string, whereString string, value interface{}, dest ...interface{}) error {
	rows, err := db.Query(`SELECT `+selectString+` FROM users WHERE `+whereString, value)
	if err != nil {
		log.Fatal(err)
		return err
	}
	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(dest...)
		if err != nil {
			log.Print(err)
			return err
		}
		return nil
	}
	return errors.New("nothing found")
}
