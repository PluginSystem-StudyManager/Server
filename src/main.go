package main

//go:generate go generate server/marketplace
//go:generate go generate server/login
//go:generate go generate server/profile
//go:generate go generate server/register
//go:generate go generate server/homepage
//go:generate go generate server/plugins
//go:generate go generate server/downloadApplication
//go:generate go generate server/db
//go:generate go generate server/dev_guide

import (
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
	"server/db"
	"server/dev_guide"
	"server/downloadApplication"
	"server/homepage"
	"server/login"
	"server/marketplace"
	"server/plugins"
	"server/profile"
	"server/register"
	"server/server"
	"server/utils"
)

func main() {
	router := httprouter.New()
	router.GET("/dist/*filepath", serveStatic)

	router.GET("/test", func(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
		writer.Write([]byte("Test"))
	})
	homepage.Init(router)
	login.Init(router)
	register.Init(router)
	plugins.Init(router)
	marketplace.Init(router)
	profile.Init(router)
	downloadApplication.Init(router)
	dev_guide.Init(router)

	db.Init()
	defer db.Close()

	s := server.New(":8080", router)
	err := s.ListenAndServe()
	if err != nil {
		log.Fatalf("Server error: %v", err)
	}

}

func serveStatic(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	http.ServeFile(w, r, utils.StaticFile(p.ByName("filepath")))
}
