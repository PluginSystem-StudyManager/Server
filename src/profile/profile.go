package profile

//go:generate jade -pkg=views -writer -d ../views profile.jade

import (
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
	"server/db"
	"server/plugins"
	"server/views"
	"server/web_lib"
)

func Init(router *httprouter.Router) {
	router.GET("/profile", profile)
}

func profile(writer http.ResponseWriter, request *http.Request, _ httprouter.Params) {
	header := web_lib.BuildHeaderData(request)
	username := header.UserName
	email := "John@gmx.de"
	token, err := db.PermanentTokenByUsername(username)
	if err != nil {
		log.Println("Error: Token should always exist!")
		token = "Not set"
	}
	pluginsData := plugins.DbDataToTemplateData(db.ListPluginsByUser(username))
	views.Profile(header, username, email, token, pluginsData, writer)
}
