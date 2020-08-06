package db

import (
	"database/sql"
	"errors"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"os"
)

var db *sql.DB

func Init() {
	dbPath := "../dist/foo.db"
	os.Remove(dbPath)

	dbLoc, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		log.Fatal(err)
		return
	}
	db = dbLoc

	// CREATE Tables
	stmt := `
	CREATE TABLE users (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		username TEXT NOT NULL UNIQUE ,
		password TEXT NOT NULL ,
		token TEXT UNIQUE ,
		token_ttl TEXT
	);
	
	CREATE TABLE user_plugins (
		user INTEGER NOT NULL REFERENCES users,
		plugin INTEGER NOT NULL REFERENCES plugins,
		PRIMARY KEY (user, plugin)
	);
	
	CREATE TABLE plugins (
		id INTEGER PRIMARY KEY AUTOINCREMENT ,
		name TEXT NOT NULL UNIQUE ,
		shortDescription TEXT NOT NULL
	);
	
	CREATE TABLE plugins_tags (
		plugin INTEGER NOT NULL REFERENCES plugins,
		tag TEXT NOT NULL --TODO: Extra table?
	);
	`
	_, err = db.Exec(stmt)
	if err != nil {
		log.Fatal(err)
		return
	}
	// TODO: Debug only
	_, _ = AddUser("John", "12345")
	_, _ = UpdateToken("John", "12345", "2022")
}

func Close() {
	_ = db.Close()
}

// Users

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

type PluginData struct {
	Name             string
	ShortDescription string
	Tags             []string
	UserIds          []int
}

func AddPlugin(data PluginData) error {
	res, err := insert("INSERT INTO plugins(Name, shortDescription) values(?, ?)", data.Name, data.ShortDescription)
	if err != nil {
		log.Fatal(err)
		return err
	}
	pluginId, err := res.LastInsertId()
	for _, tag := range data.Tags {
		res, err = insert("INSERT INTO plugins_tags(plugin, tag) values(?, ?)", pluginId, tag)
		if err != nil {
			log.Fatal(err)
			return err
		}
	}
	for _, userId := range data.UserIds {
		res, err = insert("INSERT INTO user_plugins(user, plugin) values(?, ?)", userId, pluginId)
		if err != nil {
			log.Fatal(err)
			return err
		}
	}
	return nil
}

// Utils

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
