package homepage

//go:generate jade -pkg=views -writer -d ../views home.jade

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	"server/plugins"
	"server/views"
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
	views.Homepage(pluginsTemplateData, writer)
}
