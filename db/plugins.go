package main

//go:generate schema-generate  -o plugins.schema.go -p main ../schemas/plugins/pluginData.schema.json ../schemas/plugins/list.schema.json ../schemas/plugins/add.schema.json

import (
	"encoding/json"
	"io/ioutil"
	"lang.yottadb.com/go/yottadb"
	"log"
	"net/http"
	"strconv"
	"strings"
)

const (
	cPlugins          = "^plugins"
	cName             = "name"
	cShortDescription = "shortDescription"
	cTags             = "tags"
	cAuthors          = "authors"
)

func addPlugin(w http.ResponseWriter, r *http.Request) {
	var data PluginData
	body, _ := ioutil.ReadAll(r.Body)
	_ = json.Unmarshal(body, &data)
	err := addPluginImpl(data)
	var res AddResult
	if err != nil {
		res = AddResult{
			Success: false,
			Message: err.Error(),
		}
	} else {
		res = AddResult{
			Success: true,
			Message: "",
		}
	}
	response, _ := json.Marshal(res)
	w.Write(response)
}

func addPluginImpl(data PluginData) error {
	// TODO: better handling of errors. Not just return and keep some data in db
	err := yottadb.SetValE(yottadb.NOTTP, nil, data.Name, cPlugins, []string{data.Name, cName})
	if err != nil {
		return err
	}
	err = yottadb.SetValE(yottadb.NOTTP, nil, data.ShortDescription, cPlugins, []string{data.Name, cShortDescription})
	if err != nil {
		return err
	}
	for i, tag := range data.Tags {
		err = yottadb.SetValE(yottadb.NOTTP, nil, tag, cPlugins, []string{data.Name, cTags, strconv.Itoa(i)})
		if err != nil {
			return err
		}
	}
	for i, author := range data.UserIds {
		err = yottadb.SetValE(yottadb.NOTTP, nil, strconv.Itoa(author), cPlugins, []string{data.Name, cAuthors, strconv.Itoa(i)})
		if err != nil {
			return err
		}
	}
	return nil
}

func listPlugins(w http.ResponseWriter, req *http.Request) {
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		// TODO: handle
	}
	var data ListRequest
	_ = json.Unmarshal(body, &data)
	plugins, err := listPluginsImpl(data)
	var res ListResult
	if err != nil {
		res = ListResult{
			Data:    nil,
			Message: err.Error(),
			Success: false,
		}
	} else {
		res = ListResult{
			Data:    plugins,
			Message: "",
			Success: true,
		}
	}
	response, err := json.Marshal(res)
	if err != nil {
		// TODO: handle
	}
	_, _ = w.Write(response)
}

func listPluginsImpl(request ListRequest) ([]*PluginData, error) {
	search := request.Search
	userId := request.UserId

	var pluginName = ""
	var plugins []*PluginData
	var err error
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
		log.Printf("PluginName: %s\n", pluginName)
		if len(pluginName) > 0 {
			// description
			description, _ := yottadb.ValE(yottadb.NOTTP, nil, cPlugins, []string{pluginName, cShortDescription})
			// name
			name, _ := yottadb.ValE(yottadb.NOTTP, nil, cPlugins, []string{pluginName, cName})
			// User ids
			var userIds []int
			var val = ""
			for true {
				val, err = yottadb.SubNextE(yottadb.NOTTP, nil, cPlugins, []string{pluginName, cAuthors, val})
				if err != nil {
					errorCode := yottadb.ErrorCode(err)
					if errorCode == yottadb.YDB_ERR_NODEEND {
						break
					} else {
						log.Printf("Error reading author ids: %v", err)
						// TODO: handle
					}
				}
				if len(val) > 0 {
					idString, err := yottadb.ValE(yottadb.NOTTP, nil, cPlugins, []string{pluginName, cAuthors, val})

					id, err := strconv.Atoi(idString)
					if err != nil {
						log.Printf("Error converting string to int: %v, %v", val, err)
						// TODO: handle
					}
					log.Printf("AuthorId: %v\n", id)
					userIds = append(userIds, id)
				}
			}
			// tags
			var tags []string
			val = ""
			for true {
				val, err = yottadb.SubNextE(yottadb.NOTTP, nil, cPlugins, []string{pluginName, cTags, val})
				if err != nil {
					errorCode := yottadb.ErrorCode(err)
					if errorCode == yottadb.YDB_ERR_NODEEND {
						break
					} else {
						log.Printf("Error reading tags: %v", err)
						// TODO: handle
					}
				}
				if len(val) > 0 {
					tag, err := yottadb.ValE(yottadb.NOTTP, nil, cPlugins, []string{pluginName, cTags, val})
					if err != nil {
						log.Printf("Error getting value: %v", err)
						// TODO: handle
					}
					tags = append(tags, tag)
				}
			}

			plugin := PluginData{
				Name:             name,
				ShortDescription: description,
				Tags:             tags,
				UserIds:          userIds,
			}
			if hasSearch(plugin, search, userId) {
				plugins = append(plugins, &plugin)
			}
		}
	}
	return plugins, nil
}

func hasSearch(plugin PluginData, search string, userId int) bool {
	contains := func(s string, substr string) bool {
		return strings.Contains(strings.ToUpper(s), strings.ToUpper(substr))
	}
	var idCheck bool
	if userId < 0 {
		idCheck = true
	} else {
		idCheck = false
		for _, id := range plugin.UserIds {
			if id == userId {
				idCheck = true
				break
			}
		}
	}
	return (contains(plugin.Name, search) || contains(plugin.ShortDescription, search)) && idCheck
}
