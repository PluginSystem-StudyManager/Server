package register

//go:generate jade -pkg=views -writer -d ../views register.jade

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	"server/db"
	"server/views"
)

func Init(router *httprouter.Router) {
	router.GET("/register", register)
	router.POST("/checkUserName", checkUserName)
}

func register(writer http.ResponseWriter, request *http.Request, _ httprouter.Params) {
	views.Register(writer)
}

func checkUserName(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {

	err := request.ParseForm()
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	userName := request.Form.Get("UserName")
	password := request.Form.Get("Password")
	firstName := request.Form.Get("FirstName")
	lastName := request.Form.Get("LastName")
	eMail := request.Form.Get("EMail")

	result, err := db.UsernameAvailable(userName)

	if result {
		err := db.AddUser(userName, password, firstName, lastName, eMail)
		if err != nil {

			//do Fehlerbeseitigung
		}
	} else {

		//
	}

}
