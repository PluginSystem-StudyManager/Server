package marketplace

//go:generate jade -pkg=marketplace -writer hello.jade

import (
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
	"server/db"
)

func Init(router *httprouter.Router) {
	router.GET("/marketplace", marketplace)
}

type pluginTemplate struct {
	Name             string
	ShortDescription string
	Preview          string
}

type allPluginsTemplate struct {
	Plugins string
}

func marketplace(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	_, err := db.ListPlugins()
	if err != nil {
		log.Printf("%v", err)
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte(err.Error()))
	}
	Index("Hello", "Hello", w)
}
