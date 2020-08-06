package marketplace

import (
	"bytes"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
	"server/db"
	"text/template"
)

func Init(router *httprouter.Router) {
	router.GET("/marketplace", marketplace)
}

func marketplace(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	plugins, err := db.ListPlugins()
	if err != nil {
		log.Printf("%v", err)
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte(err.Error()))
	}

	var t *template.Template
	t, err = template.ParseFiles("marketplace/template_pluginPreview.html", "marketplace/template_allPluginsPreview.html")
	if err != nil {
		log.Fatal(err)
	}
	var buffer bytes.Buffer
	for _, plugin := range plugins {
		err = t.ExecuteTemplate(&buffer, "template_pluginPreview.html", plugin)
		if err != nil {
			log.Fatal(err)
		}
	}
	data := struct {
		Plugins string
	}{Plugins: string(buffer.Bytes())}
	var allPluginsBuffer bytes.Buffer
	err = t.ExecuteTemplate(&allPluginsBuffer, "template_allPluginsPreview.html", data)
	if err != nil {
		log.Fatal(err)
	}
	_, _ = w.Write(allPluginsBuffer.Bytes())
}
