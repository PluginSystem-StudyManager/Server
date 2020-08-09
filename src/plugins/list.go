package plugins

import (
	"encoding/json"
	"fmt"
	"github.com/julienschmidt/httprouter"
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

func DbDataToTemplateData(data []db.PluginData, err error) []views.PluginTemplateData {
	if err != nil {
		return []views.PluginTemplateData{}
	}
	var pluginsTemplateData []views.PluginTemplateData
	for _, plugin := range data {
		pluginsTemplateData = append(pluginsTemplateData, views.PluginTemplateData{
			Name:             plugin.Name,
			ShortDescription: plugin.ShortDescription,
		})
	}
	return pluginsTemplateData
}
