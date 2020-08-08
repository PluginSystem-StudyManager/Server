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

func marketplace(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	plugins, err := db.ListPlugins()
	if err != nil {
		log.Printf("%v", err)
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte(err.Error()))
	}
	var pluginsTemplateDate []views.PluginTemplateData
	for _, plugin := range plugins {
		pluginsTemplateDate = append(pluginsTemplateDate, views.PluginTemplateData{
			Name:             plugin.Name,
			ShortDescription: plugin.ShortDescription,
		})
	}

	views.Marketplace(pluginsTemplateDate, w)
}
