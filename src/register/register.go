package register

//go:generate jade -pkg=views -writer -d ../views register.jade

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	"server/views"
)

func Init(router *httprouter.Router) {
	router.GET("/register", register)
}

func register(writer http.ResponseWriter, request *http.Request, _ httprouter.Params) {
	views.Register(writer)
}
