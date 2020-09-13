package register

//go:generate jade -pkg=views -writer -d ../views register.jade

import (
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"server/db"
	"server/utils"
	"server/views"
	"server/web_lib"
)

func Init(router *httprouter.Router) {
	router.GET("/register", register)
	router.POST("/checkUserName", checkUserName) // TODO: rename function. also in ts
}

func register(writer http.ResponseWriter, request *http.Request, _ httprouter.Params) {
	header := web_lib.BuildHeaderData(request)
	views.Register(header, writer)
}

func checkUserName(writer http.ResponseWriter, request *http.Request, _ httprouter.Params) {

	err := request.ParseForm()
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	userName := request.Form.Get("UserName")
	password := request.Form.Get("Password")
	eMail := request.Form.Get("EMail")

	result, err := db.UsernameAvailable(userName)

	type ErrorMessage struct {
		Fehlercode    int
		Fehlermeldung string
	}

	if result {
		passwordHash, err := utils.HashPassword(password)
		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
			return
		}
		err = db.AddUser(userName, passwordHash, eMail)
		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
			return
		}

		errorM := ErrorMessage{0, "No Error"}
		js, err := json.Marshal(errorM)
		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError) // TODO: undo register?
			return
		}
		web_lib.CreateCookie(writer, userName)
		writer.Header().Set("Content-Type", "application/json")
		_, _ = writer.Write(js)

	} else {

		errorM := ErrorMessage{5, "Dieser User existiert bereits"}
		js, err := json.Marshal(errorM)
		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
			return
		}

		writer.Header().Set("Content-Type", "application/json")
		_, _ = writer.Write(js)
	}

}
