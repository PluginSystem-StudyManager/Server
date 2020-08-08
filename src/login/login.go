package login

//go:generate jade -pkg=views -writer -d ../views login.jade

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	"server/views"
)

func Init(router *httprouter.Router) {
	router.GET("/login", login)
}

func login(writer http.ResponseWriter, request *http.Request, _ httprouter.Params) {
	views.Login("Login", writer)
}
