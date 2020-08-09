package login

//go:generate jade -pkg=views -writer -d ../views login.jade

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"server/views"
)

func Init(router *httprouter.Router) {
	router.GET("/login", login)
	router.POST("/userLogin", userLogin)
}

func login(writer http.ResponseWriter, request *http.Request, _ httprouter.Params) {
	views.Login(writer)
}

func userLogin(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {

	fmt.Printf("Funktion wird aufgerufen\n")

	request.ParseForm()
	user := request.Form.Get("userName")
	pw := request.Form.Get("UserPassword")

	fmt.Fprintf(writer, "User = %s\n", user)
	fmt.Fprintf(writer, "Passwort = %s\n", pw)

	fmt.Println("username:", request.Form["userName"])
	fmt.Println("password:", request.Form["UserPassword"])

}
