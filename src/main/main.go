package main

import (
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
	"os"
	"time"
)

func initAllJsonFile() {
	if _, err := os.Stat(allTabsFilePath); os.IsNotExist(err) {
		initValue := AllTabs{
			Timestamp: time.Now().String(),
			Plugins:   make(map[string]TabInfo),
		}
		err = WriteJson(allTabsFilePath, initValue)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func resetAllJsonFile() {
	if _, err := os.Stat(allTabsFilePath); !os.IsNotExist(err) {
		err = os.Remove(allTabsFilePath)
		if err != nil {
			log.Fatal(err)
		}
	}
	initAllJsonFile()
}

func initFolders() {
	paths := []string{pluginsPath, pluginsTmp}
	for _, path := range paths {
		if _, err := os.Stat(path); os.IsNotExist(err) {
			err := os.MkdirAll(path, os.ModeDir)
			if err != nil {
				log.Fatal(err)
			}
		}
	}

}

type TabInfoTemplate struct {
	Name string
	Body string
}

func main() {
	initFolders()
	initAllJsonFile()

	router := httprouter.New()
	AddWebsiteRoutes(router)
	AddPluginRoutes(router)

	log.Println("Start listening on: http://127.0.0.1:8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
