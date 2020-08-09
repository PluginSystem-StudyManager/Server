package marketplace

//go:generate jade -pkg=views -writer -d ../views marketplace.jade

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	"server/db"
	"server/plugins"
	"server/views"
)

func Init(router *httprouter.Router) {
	router.GET("/marketplace", marketplace)
}

func marketplace(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	_ = r.ParseForm()
	search := r.Form.Get("search")
	pluginsTemplateData := plugins.DbDataToTemplateData(db.ListPluginsSearch(search))
	views.Marketplace(pluginsTemplateData, search, w)
}
