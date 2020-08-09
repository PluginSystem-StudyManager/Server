package downloadApplication

//go:generate jade -pkg=views -writer -d ../views downloadApplication.jade

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	"server/views"
)

func Init(router *httprouter.Router) {
	router.GET("/downloadApplication", login)
}

func login(writer http.ResponseWriter, request *http.Request, _ httprouter.Params) {
	views.DownloadApplication(writer)
}
