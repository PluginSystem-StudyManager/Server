package plugins

import (
	"encoding/json"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
	"server/db"
	"server/views"
)

func list(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	plugins, err := db.ListPlugins()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte(fmt.Sprintf("error getting plugins: %v", err)))
	}
	jsonData, err := json.Marshal(plugins)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte(fmt.Sprintf("error getting plugins: %v", err)))
	}
	w.Header().Set("Content-Type", "application/json")
	_, _ = w.Write(jsonData)
}

func ListTemplateData(search string) ([]views.PluginTemplateData, error) {
	var plugins []db.PluginData
	var err error
	if len(search) > 0 {
		plugins, err = db.ListPluginsSearch(search)

	} else {
		plugins, err = db.ListPlugins()
	}
	if err != nil {
		log.Printf("%v", err)
		return nil, err
	}

	var pluginsTemplateData []views.PluginTemplateData
	for _, plugin := range plugins {
		pluginsTemplateData = append(pluginsTemplateData, views.PluginTemplateData{
			Name:             plugin.Name,
			ShortDescription: plugin.ShortDescription,
		})
	}
	return pluginsTemplateData, nil
}
