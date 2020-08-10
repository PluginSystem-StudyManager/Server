package main

import (
	"encoding/json"
	"io/ioutil"
	"lang.yottadb.com/go/yottadb"
	"log"
	"net/http"
)

type PluginData struct {
	Name             string   `json:"name"`
	ShortDescription string   `json:"short_description"`
	Tags             []string `json:"tags"`
	UserIds          []int    `json:"user_ids"`
}

const (
	cPlugins          = "plugins"
	cName             = "name"
	cShortDescription = "shortDescription"
	cTags             = "tags"
	cAuthors          = "authors"
)

type ListResult struct {
	Success bool         `json:"success"`
	Message string       `json:"message"`
	Data    []PluginData `json:"data"`
}

type AddResult struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

type ListRequest struct {
	Search string `json:"search"`
}

func addPlugin(w http.ResponseWriter, r *http.Request) {
	var data PluginData
	body, _ := ioutil.ReadAll(r.Body)
	_ = json.Unmarshal(body, &data)
	yottadb.SetValE(yottadb.NOTTP, nil, data.Name, cPlugins, []string{data.Name, cName})
	yottadb.SetValE(yottadb.NOTTP, nil, data.ShortDescription, cPlugins, []string{data.Name, cShortDescription})
	for i, tag := range data.Tags {
		yottadb.SetValE(yottadb.NOTTP, nil, tag, cPlugins, []string{data.Name, cTags, string(rune(i))})
	}
	for i, author := range data.UserIds {
		yottadb.SetValE(yottadb.NOTTP, nil, string(rune(author)), cPlugins, []string{data.Name, cAuthors, string(rune(i))})
	}
	response, _ := json.Marshal(AddResult{
		Success: true,
		Message: "",
	})
	w.Write(response)
}

func listPlugins(w http.ResponseWriter, req *http.Request) {
	body, err := ioutil.ReadAll(req.Body)
	var data ListRequest
	_ = json.Unmarshal(body, &data)

	var currSub = ""
	var dummyRes = ""
	for true {
		currSub, err = yottadb.SubNextE(yottadb.NOTTP, nil, "plugins", []string{currSub})
		if err != nil {
			errorCode := yottadb.ErrorCode(err)
			if errorCode == yottadb.YDB_ERR_NODEEND {
				break
			} else {
				panic(err) // TODO
			}
		}
		dummyRes += ";" + currSub
		log.Println(currSub)
	}
	plugins := []PluginData{{
		Name:             dummyRes,
		ShortDescription: "",
		Tags:             nil,
		UserIds:          nil,
	}}
	response, err := json.Marshal(ListResult{
		Success: true,
		Message: "",
		Data:    plugins,
	})
	w.Write(response)
}
