package login

//go:generate jade -pkg=views -writer -d ../views login.jade

import (
	"encoding/base64"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"server/db"
	"server/utils"
	"server/views"
	"server/web_lib"
	"time"
)

func Init(router *httprouter.Router) {
	router.GET("/login", login)
	router.GET("/logout", logout)
	router.POST("/userLogin", userLogin)
}

func login(writer http.ResponseWriter, request *http.Request, _ httprouter.Params) {
	header := web_lib.BuildHeaderData(request)
	views.Login(header, writer)
}

func logout(writer http.ResponseWriter, request *http.Request, _ httprouter.Params) {
	cookie := &http.Cookie{
		Name:    web_lib.CookieName,
		Value:   "",
		Path:    "/",
		Expires: time.Unix(0, 0),
	}
	http.SetCookie(writer, cookie)
	http.Redirect(writer, request, "/", http.StatusSeeOther)
}

func userLogin(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	err := request.ParseForm()
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	user := request.Form.Get("username")
	pw := request.Form.Get("password")

	success, err := db.CheckCredentials(user, pw)

	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	if !success {
		writer.WriteHeader(http.StatusUnauthorized)
	} else {
		createCookie(writer, user)
		writer.WriteHeader(http.StatusOK)
	}
}

func createCookie(writer http.ResponseWriter, username string) {
	token := utils.CreateToken()
	cookieValue := base64.StdEncoding.EncodeToString([]byte(token))

	ttl, err := time.ParseDuration("12h")
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}
	expire := time.Now().Add(ttl)

	cookie := http.Cookie{
		Name:     web_lib.CookieName,
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
