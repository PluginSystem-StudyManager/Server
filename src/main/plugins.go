package main

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"gitlab.com/golang-commonmark/markdown"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"text/template"
	"time"
)

var pluginsPath = "public/plugins"
var allTabsFilePath = "public/plugins/all_tabs.json"
var pluginsTmp = filepath.Join(pluginsPath, "tmp_upload")

type AllTabs struct {
	Timestamp string
	Plugins   map[string]TabInfo
}

type TabInfo struct {
	Name        string
	Tooltip     string
	Description string
	Authors     []string
	Repository  string
	Version     string
}

func AddPluginRoutes(router *httprouter.Router) {
	router.GET("/plugins/download/:pluginName", handlerPluginsDownload)
	router.GET("/plugins/list/", handlerPluginsList)
	router.GET("/plugins/info/:pluginName/*resource", handlerPluginInfo)
	router.POST("/plugins/upload", handlerPluginUpload)
}

func handlerPluginsDownload(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	pluginName := params.ByName("pluginName")
	filePath := filepath.Join(pluginsPath, pluginName, pluginName+".jar")
	if _, err := os.Stat(filePath); !os.IsNotExist(err) {
		http.ServeFile(w, r, filePath)
	} else {
		w.WriteHeader(http.StatusNotFound)
		http.ServeFile(w, r, "./public/404_plugin.html")
	}
}

func handlerPluginsList(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	http.ServeFile(w, r, allTabsFilePath)
}

func handlerPluginUpload(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	log.Println("POST: Fileupload")
	// max 10 MB
	_ = r.ParseMultipartForm(10 << 20)
	file := r.MultipartForm.File["file"][0]
	fileHandle, err := file.Open()
	if err != nil {
		fmt.Print("Error opening file")
		return
	}
	content := make([]byte, file.Size)
	_, _ = fileHandle.Read(content)
	_ = fileHandle.Close()

	pluginName := r.Form.Get("name")

	zipPath := filepath.Join(pluginsTmp, fmt.Sprintf("%s.zip", pluginName))

	_ = ioutil.WriteFile(zipPath, content, os.ModePerm)
	zipTabUploaded(zipPath, pluginName)
	_ = os.Remove(zipPath)
	_, _ = fmt.Fprint(w, "Successfully uploaded file")
	log.Printf("Successfully uploaded plugin: %s", pluginName)
}

func handlerPluginInfo(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	pluginName := params.ByName("pluginName")
	resourceName := params.ByName("resource")
	if strings.HasPrefix(resourceName, "/") {
		resourceName = strings.TrimLeft(resourceName, "/")
	}
	if resourceName == "html" {
		md := markdown.New(markdown.XHTMLOutput(true))
		fileData, err := ioutil.ReadFile(filepath.Join(pluginsPath, pluginName, "info/README.md"))
		if err != nil {
			log.Fatal(err)
			return
		}
		body := md.RenderToString(fileData)
		data := TabInfoTemplate{
			Name: pluginName,
			Body: body,
		}
		var t *template.Template
		t, err = template.ParseFiles("public/templates/tab_info.html")
		if err != nil {
			log.Fatal(err)
		}
		err = t.ExecuteTemplate(w, "tab_info.html", data)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		getResource := func(paths ...string) string {
			for _, path := range paths {
				if _, err := os.Stat(path); !os.IsNotExist(err) {
					return path
				}
			}
			return ""
		}
		path := getResource(filepath.Join(pluginsPath, pluginName, "info", resourceName),
			filepath.Join("public/res", resourceName))
		if path == "" {
			w.WriteHeader(http.StatusNotFound)
			_, _ = w.Write([]byte(""))
		} else {
			http.ServeFile(w, r, path)
		}
	}
}

func zipTabUploaded(zipPath string, pluginName string) {
	// new directory with pluginName
	pluginPath := fmt.Sprintf("public/plugins/%s", pluginName)
	// copy info + jar
	if _, err := os.Stat(pluginPath); os.IsNotExist(err) {
		err := os.Mkdir(pluginPath, os.ModeDir)
		if err != nil {
			log.Fatal(err) // TODO
			return
		}
	}
	err := Unzip(zipPath, pluginPath)
	if err != nil {
		log.Fatal(err)
	}
	// add info.json to all.json
	var allTabs AllTabs
	err = UnpackJson(allTabsFilePath, &allTabs)
	if err != nil {
		log.Fatal(err)
		return
	}
	var tabInfo TabInfo
	err = UnpackJson(filepath.Join(pluginPath, "info/tab_info.json"), &tabInfo)
	if err != nil {
		log.Fatal(err)
		return
	}
	allTabs.Plugins[pluginName] = tabInfo
	// update timestamp in json
	allTabs.Timestamp = time.Now().String()

	// Write back
	err = WriteJson(allTabsFilePath, allTabs)
	if err != nil {
		log.Fatal(err)
	}
}
