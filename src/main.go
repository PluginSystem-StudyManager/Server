package main

import (
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
	"server/homepage"
	"server/server"
	"server/utils"
)

func main() {
	router := httprouter.New()
	router.GET("/dist/*filepath", serveStatic)

	homepage.Init(router)
	s := server.New(":8080", router)
	err := s.ListenAndServe()
	if err != nil {
		log.Fatalf("Server error: %v", err)
	}
}

func serveStatic(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	http.ServeFile(w, r, utils.StaticFile(p.ByName("filepath")))
}
