package db

import "lang.yottadb.com/go/yottadb"

func AddUser(username string, password string) error {
	err := yottadb.SetValE(yottadb.NOTTP, nil, username, "user", []string{username, "username"})
	err := yottadb.SetValE(yottadb.NOTTP, nil, password, "user", []string{username, "password"})
}

func CheckCredentials(username string, password string) (bool, error) {
	pw_db, err := yottadb.ValE(yottadb.NOTTP, nil, "password", "user", []string{username})
	if err != nil {
		return false, err
	}
	return pw_db == password, nil
}

func UpdateToken(username string, token string, ttl string) (sql.Result, error) {
	return nil, nil
}

func UserIdByToken(token string) (int, error) {
	return 0, nil
}
