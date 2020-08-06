package login

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func Init(router *httprouter.Router) {
	router.GET("/login", login)
}

func login(writer http.ResponseWriter, request *http.Request, _ httprouter.Params) {
	http.ServeFile(writer, request, "login/login.html")
}
