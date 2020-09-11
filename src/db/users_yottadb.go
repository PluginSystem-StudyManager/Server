//+build excluded

/*
For now only the plugins db is implemented with yottadb. Later on everything could be use yottadb.
*/

package db

func AddUser(username string, password string) error {
	return nil
}

func CheckCredentials(username string, password string) (bool, error) {
	return false, nil
}

func UpdateToken(username string, token string, ttl string) error {
	return nil
}

func UserIdByToken(token string) (int, error) {
	return 0, nil
}
