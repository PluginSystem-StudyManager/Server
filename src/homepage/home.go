package homepage

//go:generate jade -pkg=views -writer -d ../views home.jade

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	"server/views"
)

func Init(router *httprouter.Router) {
	router.GET("/", home)
}

func home(writer http.ResponseWriter, request *http.Request, _ httprouter.Params) {
	views.Homepage(writer)
}
