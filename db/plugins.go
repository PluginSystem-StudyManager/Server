package main

//go:generate schema-generate  -o plugins.schema.go -p main ../schemas/plugins/pluginData.schema.json ../schemas/plugins/list.schema.json ../schemas/plugins/add.schema.json

import (
	"encoding/json"
	"io/ioutil"
	"lang.yottadb.com/go/yottadb"
	"net/http"
	"strings"
)

const (
	cPlugins          = "plugins"
	cName             = "name"
	cShortDescription = "shortDescription"
	cTags             = "tags"
	cAuthors          = "authors"
)

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
	search := data.Search

	var pluginName = ""
	var plugins []*PluginData
	for true {
		pluginName, err = yottadb.SubNextE(yottadb.NOTTP, nil, cPlugins, []string{pluginName})
		if err != nil {
			errorCode := yottadb.ErrorCode(err)
			if errorCode == yottadb.YDB_ERR_NODEEND {
				break
			} else {
				panic(err) // TODO
			}
		}
		if len(pluginName) > 0 {
			description, _ := yottadb.ValE(yottadb.NOTTP, nil, cPlugins, []string{pluginName, cShortDescription})
			name, _ := yottadb.ValE(yottadb.NOTTP, nil, cPlugins, []string{pluginName, cName})
			plugin := PluginData{
				Name:             name,
				ShortDescription: description,
				Tags:             nil,
				UserIds:          nil,
			}
			if hasSearch(plugin, search) {
				plugins = append(plugins, &plugin)
			}
		}
	}
	response, err := json.Marshal(ListResult{
		Success: true,
		Message: "",
		Data:    plugins,
	})
	w.Write(response)
}

func hasSearch(plugin PluginData, search string) bool {
	contains := func(s string, substr string) bool {
		return strings.Contains(strings.ToUpper(s), strings.ToUpper(substr))
	}
	return contains(plugin.Name, search) || contains(plugin.ShortDescription, search)
}
