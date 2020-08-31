package web_lib

import (
	"encoding/base64"
	"encoding/json"
	"net/http"
	"server/db"
	"server/views"
)

type UserData struct {
	Token string
}

func BuildHeaderData(r *http.Request) views.HeaderData {
	cookie, err := r.Cookie("userdata") // TODO: constant cookie name
	notLoggedIn := func() views.HeaderData {
		return views.HeaderData{
			UserName: "",
			LoggedIn: false,
		}
	}
	if err != nil {
		// Can't find cookie
		return notLoggedIn()
	}
	data, err := base64.StdEncoding.DecodeString(cookie.Value)
	if err != nil {
		// Wrong formatted Cookie
		return notLoggedIn()
	}
	var userData UserData
	err = json.Unmarshal(data, &userData)
	if err != nil {
		// Wrong formatted cookie
		return notLoggedIn()
	}
	user, err := db.UserByToken(userData.Token)
	if err != nil {
		// Token does not exist or is expired. TODO: Maybe delete cookie
		return notLoggedIn()
	}
	return views.HeaderData{
		UserName: user.Username,
		LoggedIn: true,
	}
}
