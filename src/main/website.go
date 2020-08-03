package main

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	http.ServeFile(w, r, "public/website/index.html")
}

func Hello(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	fmt.Fprintf(w, "Hello %s\n", params.ByName("name"))
}

func AddWebsiteRoutes(router *httprouter.Router) {
	router.GET("/", Index)
	router.GET("/hello/:name", Hello)
}
