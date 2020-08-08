package profile

//go:generate jade -pkg=views -writer -d ../views profile.jade

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	"server/views"
)

func Init(router *httprouter.Router) {
	router.GET("/profile", register)
}

func register(writer http.ResponseWriter, request *http.Request, _ httprouter.Params) {

	name := "Hans Wurst"
	views.Profile(name, writer)
}
