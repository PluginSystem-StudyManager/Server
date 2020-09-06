package web_lib

import (
	"encoding/base64"
	"net/http"
	"server/db"
	"server/utils"
	"server/views"
	"time"
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
			EMail:    "",
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

	//TODO hier fehlt noch Mehtode "Email By Tokken.."

	if err != nil {
		// Token does not exist or is expired. TODO: Maybe delete cookie
		return notLoggedIn()
	}
	return views.HeaderData{
		UserName: user.Username,
		EMail:    "", //return der methode E-Mail ...
		LoggedIn: true,
	}
}

func CreateCookie(writer http.ResponseWriter, username string) {
	token := utils.CreateToken()
	cookieValue := base64.StdEncoding.EncodeToString([]byte(token))

	ttl, err := time.ParseDuration("12h")
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}
	expire := time.Now().Add(ttl)

	cookie := http.Cookie{
		Name:     CookieName,
		Value:    cookieValue,
		Expires:  expire,
		SameSite: http.SameSiteStrictMode, // TODO Vor dem Livebetrieb nur noch https zulassen
	}
	http.SetCookie(writer, &cookie)

	err = db.UpdateToken(username, token, expire.String())

	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError) //TODO Anpassen
		return
	}
}
