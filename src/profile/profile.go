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
	token, err := web_lib.GetUserToken(request)
	if err != nil {
		// No Token
		// TODO: this should never happen, but handle
		log.Println("ERROR: Token does not exist")
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}
	user, err := db.UserByToken(token)
	if err != nil {
		// TODO: this should never happen, but handle
		log.Printf("ERROR: username with token (%s) does not exist", token)
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}
	username := user.Username
	email := user.Email
	permanentToken, err := db.PermanentTokenByUsername(username)
	if err != nil {
		// TODO: handle
		log.Println("Error: Token should always exist!")
		token = "Not set"
	}
	pluginsData := plugins.DbDataToTemplateData(db.ListPluginsByUser(username))
	views.Profile(header, username, email, permanentToken, pluginsData, writer)
}
