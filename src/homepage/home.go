package homepage

//go:generate jade -pkg=views -writer -d ../views home.jade

import (
	"encoding/base64"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"server/db"
	"server/plugins"
	"server/views"
)

const (
	numPluginsPreview = 4 // How many plugins are shown in the homepage
)

func Init(router *httprouter.Router) {
	router.GET("/", Home)
}

func Home(writer http.ResponseWriter, request *http.Request, _ httprouter.Params) {
	pluginsTemplateData := plugins.DbDataToTemplateData(db.ListPlugins())

	locNumPlugins := numPluginsPreview
	if locNumPlugins > len(pluginsTemplateData) {
		locNumPlugins = len(pluginsTemplateData)
	}
	header := buildHeaderData(request)
	views.Homepage(header, pluginsTemplateData[:locNumPlugins], writer)
}

// TODO: Move to web-lib and use in all views
type UserDate struct {
	Token string	
}
// TODO: Move to web-lib and use in all views
func buildHeaderData(r *http.Request) views.HeaderData {
	cookie, err := r.Cookie("userdata") 	// TODO: constant cookie name
	notLoggedIn := function() views.HeaderData {
		return views.HeaderData {
			Username ""
			LoggedIn false
		}
	}
	if err != nil {
		// Can't find cookie
		return notLoggedIn()
	}
	data, err := base64.StdEncoding.DecodeString(cookie.Value)
	if err != nil {
		// Wrong formatted Cookie
		return notLoggedIn()	
	}
	var userData UserData
	err = json.Unmarshal(data, &userData)
	if err != nil {
		// Wrong formatted cookie
		return notLoggedIn()
	}
	user, err := db.UserByToken(userdata.Token) // TODO: implement in DB
	if err != nil {
		// Token does not exist or is expired. TODO: Maybe delete cookie
		return notLoggedIn()	
	}
	return views.HeaderData{
		Username user.Username
		LoggedIn true
	}
}
