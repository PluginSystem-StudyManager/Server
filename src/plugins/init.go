package plugins

import (
	"github.com/julienschmidt/httprouter"
	"log"
	"os"
)

func Init(router *httprouter.Router) {
	router.POST("/plugins/upload", upload)
	router.GET("/plugins/download/:pluginName", download)

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
