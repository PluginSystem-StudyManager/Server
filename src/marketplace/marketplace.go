package marketplace

//go:generate jade -pkg=views -writer -d ../views marketplace.jade

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	"server/plugins"
	"server/views"
)

func Init(router *httprouter.Router) {
	router.GET("/marketplace", marketplace)
}

func marketplace(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	_ = r.ParseForm()
	search := r.Form.Get("search")
	pluginsTemplateData, err := plugins.ListTemplateData(search)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		views.Marketplace([]views.PluginTemplateData{}, search, w)
		return
	}
	views.Marketplace(pluginsTemplateData, search, w)
}
