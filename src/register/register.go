package register

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	"server/utils"
)

func Init(router *httprouter.Router) {
	router.GET("/register", register)
}

func register(writer http.ResponseWriter, request *http.Request, _ httprouter.Params) {
	http.ServeFile(writer, request, utils.StaticFile("register/register.html"))
}
