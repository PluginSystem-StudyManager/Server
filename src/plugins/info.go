package plugins

//go:generate jade -pkg=views -writer -d ../views tab_info.jade
//go:generate jade -pkg=views -writer -d ../views tab_info_api.jade

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
	"server/web_lib"
	"strings"
)

func infoApi(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	pluginName := p.ByName("pluginName")
	resourceName := strings.TrimLeft(p.ByName("resource"), "/")
	switch resourceName {
	case "html":
		md := markdown.New(markdown.XHTMLOutput(true))
		fileData, err := ioutil.ReadFile(filepath.Join(pluginsPath, pluginName, "info", "README.md"))
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusNotFound)
			_, _ = w.Write([]byte("Plugin not found")) // TODO: 404 not found page
			return
		}
		body := md.RenderToString(fileData)
		views.TabInfoApi(body, pluginName, w)
	default:
		infoResource(w, r, pluginName, resourceName)
	}
}

func infoWebsite(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	pluginName := p.ByName("pluginName")
	resourceName := strings.TrimLeft(p.ByName("resource"), "/")
	switch resourceName {
	case "html":
		md := markdown.New(markdown.XHTMLOutput(true))
		fileData, err := ioutil.ReadFile(filepath.Join(pluginsPath, pluginName, "README.md"))
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusNotFound)
			_, _ = w.Write([]byte("Plugin not found")) // TODO: 404 not found page
			return
		}
		body := md.RenderToString(fileData)
		header := web_lib.BuildHeaderData(r)
		views.TabInfo(header, body, pluginName, w)
	default:
		infoResource(w, r, pluginName, resourceName)
	}
}

func infoResource(w http.ResponseWriter, r *http.Request, pluginName string, resourceName string) {
	if strings.HasPrefix(resourceName, "dist") {
		http.ServeFile(w, r, "../"+resourceName)
	} else {
		projectRes := utils.StaticFile("plugins/" + resourceName)
		pluginRes := filepath.Join(pluginsPath, pluginName, "info", resourceName)
		if _, err := os.Stat(projectRes); !os.IsNotExist(err) {
			http.ServeFile(w, r, projectRes)
		} else {
			http.ServeFile(w, r, pluginRes)
		}
	}
}
