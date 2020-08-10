package main

//go:generate go generate server/marketplace
//go:generate go generate server/login
//go:generate go generate server/profile
//go:generate go generate server/register
//go:generate go generate server/homepage
//go:generate go generate server/plugins
//go:generate go generate server/downloadApplication

import (
	"bytes"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"io/ioutil"
	"log"
	"net/http"
	"server/db"
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

	db.Init()
	defer db.Close()

	// TODO: remove: test only
	testDb()

	s := server.New(":8080", router)
	err := s.ListenAndServe()
	if err != nil {
		log.Fatalf("Server error: %v", err)
	}

}

func testDb() {
	url := "http://db:8090"
	resp, err := http.Get(url + "/hello")
	if err != nil {
		fmt.Printf("%v\n", err)
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	content := string(body)
	fmt.Printf("%v - %v\n", content, err)

	resp, err = http.Post(url+"/add", "text/plaintext", bytes.NewBuffer([]byte("NewData")))
	defer resp.Body.Close()
	body, err = ioutil.ReadAll(resp.Body)
	content = string(body)
	fmt.Printf("%v - %v\n", content, err)

	resp, err = http.Get(url + "/get")
	defer resp.Body.Close()
	body, err = ioutil.ReadAll(resp.Body)
	content = string(body)
	fmt.Printf("%v - %v\n", content, err)

	err = db.AddPlugin(db.PluginData{
		Name:             "Dummy_Plugin_1",
		ShortDescription: "Description",
		Tags:             []string{"Tag1", "tag2"},
		UserIds:          []int{1, 2},
	})
	fmt.Printf("add plugin: %v\n", err)

	err = db.AddPlugin(db.PluginData{
		Name:             "Dummy_Plugin_2",
		ShortDescription: "Description2",
		Tags:             []string{"Tag3", "tag4"},
		UserIds:          []int{1, 3},
	})
	fmt.Printf("add plugin: %v\n", err)

	pluginsList, err := db.ListPlugins()
	if err != nil {
		fmt.Printf("Error listing plugins: %v\n", err)
	} else {
		fmt.Printf("Plugins: %v\n", pluginsList)
		fmt.Printf("Plugin 1 name: %v\n", pluginsList[0].Name)
		fmt.Printf("Plugin 1 description: %v\n", pluginsList[0].ShortDescription)
	}
}

func serveStatic(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	http.ServeFile(w, r, utils.StaticFile(p.ByName("filepath")))
}
