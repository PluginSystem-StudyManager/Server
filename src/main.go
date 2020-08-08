package main

import (
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
	"server/db"
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

	homepage.Init(router)
	login.Init(router)
	register.Init(router)
	plugins.Init(router)
	marketplace.Init(router)
	profile.Init(router)

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
