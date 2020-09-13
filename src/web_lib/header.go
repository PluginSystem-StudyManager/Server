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
	notLoggedIn := func() views.HeaderData {
		return views.HeaderData{
			UserName: "",
			LoggedIn: false,
		}
	}

	token, err := GetUserToken(r)
	if err != nil {
		return notLoggedIn()
	}
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

func GetUserToken(r *http.Request) (string, error) {
	cookie, err := r.Cookie(CookieName)
	if err != nil {
		// Can't find cookie
		return "", err
	}
	t, err := base64.StdEncoding.DecodeString(cookie.Value)
	if err != nil {
		// Wrong formatted Cookie
		return "", err
	}
	token := string(t[:])
	return token, nil
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
