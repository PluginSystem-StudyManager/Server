package homepage

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	"server/utils"
)

func Init(router *httprouter.Router) {
	router.GET("/", home)
}

func home(writer http.ResponseWriter, request *http.Request, _ httprouter.Params) {
	http.ServeFile(writer, request, utils.StaticFile("homepage/index.html"))
}
