package homepage

//go:generate jade -pkg=views -writer -d ../views home.jade

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	"server/plugins"
	"server/views"
)

const (
	numPluginsPreview = 4 // How many plugins are shown in the homepage
)

func Init(router *httprouter.Router) {
	router.GET("/", Home)
}

func Home(writer http.ResponseWriter, request *http.Request, _ httprouter.Params) {
	pluginsTemplateData, err := plugins.ListTemplateData("")
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		views.Homepage([]views.PluginTemplateData{}, writer)
		return
	}

	locNumPlugins := numPluginsPreview
	if locNumPlugins > len(pluginsTemplateData) {
		locNumPlugins = len(pluginsTemplateData)
	}
	views.Homepage(pluginsTemplateData[:locNumPlugins], writer)
}
