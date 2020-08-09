package profile

//go:generate jade -pkg=views -writer -d ../views profile.jade

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	"server/plugins"
	"server/views"
)

func Init(router *httprouter.Router) {
	router.GET("/profile", profile)
}

func profile(writer http.ResponseWriter, request *http.Request, _ httprouter.Params) {
	pluginsData, err := plugins.ListTemplateData("")
	if err != nil {
		pluginsData = []views.PluginTemplateData{}
	}
	name := "Hans Wurst"
	views.Profile(name, pluginsData, writer)
}
