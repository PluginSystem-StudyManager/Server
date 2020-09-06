package db

import (
	"errors"
	"log"
	"server/utils"
)

func AddUser(username string, password string, email string) error {
	permanentToken := utils.CreateToken()
	_, err := insert(
		"INSERT INTO users(username, password, e_mail, permanent_token) values (?, ?, ?, ?)",
		username, password, email, permanentToken)
	return err
}

func AddDebugUser() {
	passwordHash, _ := utils.HashPassword("12345")
	_, _ = insert(
		"INSERT INTO users(username, password, e_mail, permanent_token, token, token_ttl) values (?, ?, ?, ?, ?, ?)",
		"John", passwordHash, "john@ross.com", "12345", "12345", "2099")
}

func CheckCredentials(username string, password string) (bool, error) {

	rows, err := db.Query(`SELECT password FROM users where  username=?`, username)
	if err != nil {
		log.Printf("Query Error: %v", err)
		return false, err
	}

	defer rows.Close()
	for rows.Next() {

		var dbPassword string

		err := rows.Scan(&dbPassword)
		if err != nil {
			log.Printf("Query Error: %v", err)
			return false, err
		}
		log.Println(password + " ===" + dbPassword)
		if utils.CheckPasswordHash(password, dbPassword) {
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

//func EMailByToken(token string) (string, error) {

// get User by token ....

//}

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
