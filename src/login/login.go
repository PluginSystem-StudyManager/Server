package login

//go:generate jade -pkg=views -writer -d ../views login.jade

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"math/rand"
	"net/http"
	"server/db"
	"server/views"
	"strconv"
	"time"
)

func Init(router *httprouter.Router) {
	router.GET("/login", login)
	router.POST("/userLogin", userLogin)
}

func login(writer http.ResponseWriter, request *http.Request, _ httprouter.Params) {
	views.Login(writer)
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

	min := 1000000000
	max := 3000000000

	userID := strconv.Itoa(rand.Intn(max-min) + min)
	fmt.Println(userID)

	ttl, err := time.ParseDuration("12h")
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}
	expire := time.Now().Add(ttl)

	cookie := http.Cookie{
		Name:     "UserKey",
		Value:    userID,
		Expires:  expire,
		SameSite: http.SameSiteStrictMode, // TODO Vor dem Livebetrieb nur noch https zulassen
	}
	http.SetCookie(writer, &cookie)

	err = db.UpdateToken(username, userID, ttl.String())

	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError) //TODO Anpassen
		return
	}
}
