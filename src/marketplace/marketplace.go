package marketplace

//go:generate jade -pkg=views -writer -d ../views marketplace.jade

import (
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
	"server/db"
	"server/views"
)

func Init(router *httprouter.Router) {
	router.GET("/marketplace", marketplace)
}

type allPluginsTemplate struct {
	Plugins string
}

func marketplace(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	_, err := db.ListPlugins()
	if err != nil {
		log.Printf("%v", err)
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte(err.Error()))
	}
	views.Marketplace("Marketplace", "Lukas", []views.PluginTemplate{{
		Name:             "MyPLugin",
		ShortDescription: "Some dsc",
		Preview:          "",
	},
		{
			Name:             "Anpother plugin",
			ShortDescription: "Some other description",
			Preview:          "",
		}}, w)
}
