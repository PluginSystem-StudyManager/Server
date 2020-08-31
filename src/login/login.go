package login

//go:generate jade -pkg=views -writer -d ../views login.jade

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	"server/db"
	"server/views"
	"server/web_lib"
)

func Init(router *httprouter.Router) {
	router.GET("/login", login)
	router.POST("/userLogin", userLogin)
}

func login(writer http.ResponseWriter, request *http.Request, _ httprouter.Params) {
	header := web_lib.BuildHeaderData(request)
	views.Login(header, writer)
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
		writer.WriteHeader(http.StatusOK)
	}
}
