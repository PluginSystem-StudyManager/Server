package homepage

//go:generate jade -pkg=views -writer -d ../views home.jade

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	"server/db"
	"server/plugins"
	"server/views"
	"server/web_lib"
)

const (
	numPluginsPreview = 4 // How many plugins are shown in the homepage
)

func Init(router *httprouter.Router) {
	router.GET("/", Home)
}

func Home(writer http.ResponseWriter, request *http.Request, _ httprouter.Params) {
	pluginsTemplateData := plugins.DbDataToTemplateData(db.ListPlugins())

	locNumPlugins := numPluginsPreview
	if locNumPlugins > len(pluginsTemplateData) {
		locNumPlugins = len(pluginsTemplateData)
	}
	header := web_lib.BuildHeaderData(request)
	views.Homepage(header, pluginsTemplateData[:locNumPlugins], writer)
}
