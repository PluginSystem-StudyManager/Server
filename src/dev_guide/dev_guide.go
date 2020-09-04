package dev_guide

//go:generate jade -pkg=views -writer -d ../views dev_guide.jade

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	"server/views"
	"server/web_lib"
)

func Init(router *httprouter.Router) {
	router.GET("/devguide", devGuide)
}

func devGuide(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	header := web_lib.BuildHeaderData(r)

	views.DevGuide(header, w)
}
