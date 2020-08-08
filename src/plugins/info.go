package plugins

//go:generate jade -pkg=views -writer -d ../views tab_info.jade

import (
	"github.com/julienschmidt/httprouter"
	"gitlab.com/golang-commonmark/markdown"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"server/utils"
	"server/views"
	"strings"
)

type tabInfoTemplate struct {
	Name string
	Body string
}

func info(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	pluginName := p.ByName("pluginName")
	resourceName := strings.TrimLeft(p.ByName("resource"), "/")
	switch resourceName {
	case "html":
		md := markdown.New(markdown.XHTMLOutput(true))
		fileData, err := ioutil.ReadFile(filepath.Join(pluginsPath, pluginName, "README.md"))
		if err != nil {
			log.Fatal(err)
			return
		}
		body := md.RenderToString(fileData)
		views.TabInfo(body, pluginName, w)
	default:
		if strings.HasPrefix(resourceName, "dist") {
			http.ServeFile(w, r, "../"+resourceName)
		} else {
			projectRes := utils.StaticFile("plugins/" + resourceName)
			pluginRes := filepath.Join(pluginsPath, pluginName, resourceName)
			if _, err := os.Stat(projectRes); !os.IsNotExist(err) {
				http.ServeFile(w, r, projectRes)
			} else {
				http.ServeFile(w, r, pluginRes)
			}
		}
	}
}
