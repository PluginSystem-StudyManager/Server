//+build linux

package db

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
)

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
	// TODO: implement
	return 0, errors.New("not exists")
}

func ListPlugins() ([]*PluginData, error) {
	req := ListRequest{
		Search: "",
		UserId: -1,
	}
	return listPlugins(req)
}

func ListPluginsSearch(value string) ([]*PluginData, error) {
	req := ListRequest{
		Search: value,
		UserId: -1,
	}
	return listPlugins(req)
}

func ListPluginsByUser(username string) ([]*PluginData, error) {
	id, err := UserIdByUsername(username)
	if err != nil {
		return nil, err
	}
	req := ListRequest{
		Search: "",
		UserId: id,
	}
	return listPlugins(req)
}

func listPlugins(req ListRequest) ([]*PluginData, error) {
	body, _ := json.Marshal(req)
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
