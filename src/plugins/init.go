package plugins

import (
	"github.com/julienschmidt/httprouter"
	"log"
	"os"
)

func Init(router *httprouter.Router) {
	router.GET("/plugins/info/:pluginName/*resource", infoWebsite)
	router.POST("/api/plugins/upload", upload)
	router.GET("/api/plugins/download/:pluginName", download)
	router.GET("/api/plugins/info/:pluginName/*resource", infoApi)
	router.GET("/api/plugins/list", list)

	// folders
	mkIfNotExist := func(path string) error {
		if _, err := os.Stat(path); os.IsNotExist(err) {
			return os.MkdirAll(path, os.ModeDir)
		}
		return nil
	}
	err := mkIfNotExist(pluginsPath)
	if err != nil {
		log.Fatal(err)
	}
	_ = mkIfNotExist(pluginsTmpPath)
}
