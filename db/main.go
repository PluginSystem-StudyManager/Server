package main

import (
	"fmt"
	"io/ioutil"
	"lang.yottadb.com/go/yottadb"
	"log"
	"net/http"
)

func add(w http.ResponseWriter, req *http.Request) {
	body, err := ioutil.ReadAll(req.Body)
	name := string(body)
	err = yottadb.SetValE(yottadb.NOTTP, nil, name, "username", []string{"sub"})
	if err != nil {
		fmt.Fprintf(w, "Fail")
		return
	}
	fmt.Fprintf(w, "Success")
}

func get(w http.ResponseWriter, req *http.Request) {
	r, err := yottadb.ValE(yottadb.NOTTP, nil, "username", []string{"sub"})
	if err != nil {
		fmt.Fprintf(w, "Fail")
		return
	}
	fmt.Fprintf(w, r)
}

func hello(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "Hello from DB\n")
}

func main() {
	http.HandleFunc("/hello", hello)
	http.HandleFunc("/add", add)
	http.HandleFunc("/get", get)
	http.HandleFunc("/plugins/add", addPlugin)
	http.HandleFunc("/plugins/list", listPlugins)
	log.Println("DB started at port 8090")
	http.ListenAndServe(":8090", nil)
}
