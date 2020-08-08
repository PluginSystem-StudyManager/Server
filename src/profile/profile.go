package profile

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	"server/utils"
)

func Init(router *httprouter.Router) {
	router.GET("/profile", register)
}

func register(writer http.ResponseWriter, request *http.Request, _ httprouter.Params) {
	http.ServeFile(writer, request, utils.StaticFile("profile/profile.html"))
}
