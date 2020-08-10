//+build yottadb

package db

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
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

func AddPlugin(data PluginData) error {
	body, _ := json.Marshal(data)
	res, err := http.Post("http://db:8090/plugins/add", "application/json", bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	defer res.Body.Close()
	bodyBytes, _ := ioutil.ReadAll(res.Body)
	var response AddResult
	_ = json.Unmarshal(bodyBytes, &response)
	if response.Success {
		return nil
	} else {
		return errors.New("Error adding plugin: " + response.Message)
	}
}

func PluginIdByName(name string) (int, error) { // TODO: Change name to exists()
	return 0, errors.New("not exists")
}

func ListPlugins() ([]PluginData, error) {
	return ListPluginsSearch("")
}

func ListPluginsSearch(value string) ([]PluginData, error) {
	body, _ := json.Marshal(ListRequest{
		Search: value,
	})
	res, err := http.Post("http://db:8090/plugins/list", "application/json", bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	bodyBytes, _ := ioutil.ReadAll(res.Body)
	var response ListResult
	_ = json.Unmarshal(bodyBytes, &response)
	if response.Success {
		return response.Data, nil
	} else {
		return nil, errors.New("Error listing plugin: " + response.Message)
	}
}

func ListPluginsByUser(username string) ([]PluginData, error) {
	return nil, nil
}
