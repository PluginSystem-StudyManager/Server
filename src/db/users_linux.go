package db

func AddUser(username string, password string) (sql.Result, error) {

}

func CheckCredentials(username string, password string) (bool, error) {

}

func UpdateToken(username string, token string, ttl string) (sql.Result, error) {

}

func UserIdByToken(token string) (int, error) {

}
