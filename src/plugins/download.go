package plugins

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	"os"
	"path/filepath"
)

func download(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	pluginName := p.ByName("pluginName")
	filePath := filepath.Join(pluginsPath, pluginName, "plugin_"+pluginName+".jar")
	if _, err := os.Stat(filePath); !os.IsNotExist(err) {
		http.ServeFile(w, r, filePath)
	} else {
		w.WriteHeader(http.StatusNotFound)
		_, _ = w.Write([]byte("Plugin not found")) // TODO: Server 404 html file
	}
}
