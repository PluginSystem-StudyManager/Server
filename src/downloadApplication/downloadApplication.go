package downloadApplication

//go:generate jade -pkg=views -writer -d ../views downloadApplication.jade

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	"server/views"
	"server/web_lib"
)

func Init(router *httprouter.Router) {
	router.GET("/downloadApplication", login)
}

func login(writer http.ResponseWriter, request *http.Request, _ httprouter.Params) {
	header := web_lib.BuildHeaderData(request)
	views.DownloadApplication(header, writer)
}
