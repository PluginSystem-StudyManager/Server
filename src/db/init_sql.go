//+build !linux

package db

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"os"
)

var db *sql.DB

func Init() {
	log.Println("Initialized db on Windows")
	dbPath := "../dist/foo.db"
	_ = os.Remove(dbPath)

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
		token_ttl TEXT NULL ,
		permanent_token TEXT UNIQUE,
		firstName Text NULL ,
		lastName Text NULL ,
		e_mail Text NULL
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
	_ = AddUser("John", "12345", "John", "Maier", "John.Maier@erb.de")
	_ = UpdateToken("John", "12345", "2022")
}

func Close() {
	_ = db.Close()
}
