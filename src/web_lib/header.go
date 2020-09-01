package web_lib

import (
	"encoding/base64"
	"net/http"
	"server/db"
	"server/views"
)

const CookieName = "UserToken"

type UserData struct {
	Token string
}

func BuildHeaderData(r *http.Request) views.HeaderData {
	cookie, err := r.Cookie(CookieName)
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
	t, err := base64.StdEncoding.DecodeString(cookie.Value)
	if err != nil {
		// Wrong formatted Cookie
		return notLoggedIn()
	}
	token := string(t[:])
	user, err := db.UserByToken(token)
	if err != nil {
		// Token does not exist or is expired. TODO: Maybe delete cookie
		return notLoggedIn()
	}
	return views.HeaderData{
		UserName: user.Username,
		LoggedIn: true,
	}
}
