package login

//go:generate jade -pkg=views -writer -d ../views login.jade

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"server/db"
	"server/homepage"
	"server/views"
)

func Init(router *httprouter.Router) {
	router.GET("/login", login)
	router.POST("/userLogin", userLogin)
}

func login(writer http.ResponseWriter, request *http.Request, _ httprouter.Params) {
	views.Login(writer)
}

func userLogin(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {

	fmt.Printf("Funktion wird aufgerufen\n")

	request.ParseForm()
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
		homepage.Home(writer, request, params)
	}
}
