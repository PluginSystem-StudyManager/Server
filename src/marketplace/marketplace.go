package marketplace

//go:generate jade -pkg=views -writer -d ../views marketplace.jade

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	"server/db"
	"server/plugins"
	"server/views"
	"server/web_lib"
)

func Init(router *httprouter.Router) {
	router.GET("/marketplace", marketplace)
}

func marketplace(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	_ = r.ParseForm()
	search := r.Form.Get("search")
	pluginsTemplateData := plugins.DbDataToTemplateData(db.ListPluginsSearch(search))
	header := web_lib.BuildHeaderData(r)
	views.Marketplace(header, pluginsTemplateData, search, w)
}
