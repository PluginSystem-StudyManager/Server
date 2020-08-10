//+build yottadb

package db

import (
	"lang.yottadb.com/go/yottadb"
	"log"
)

const (
	cPlugins          = "plugins"
	cName             = "name"
	cShortDescription = "shortDescription"
	cTags             = "tags"
	cAuthors          = "authors"
)

func AddPlugin(data PluginData) error {
	yottadb.SetValE(yottadb.NOTTP, nil, data.Name, cPlugins, []string{data.Name, cName})
	yottadb.SetValE(yottadb.NOTTP, nil, data.ShortDescription, cPlugins, []string{data.Name, cShortDescription})
	for i, tag := range data.Tags {
		yottadb.SetValE(yottadb.NOTTP, nil, tag, cPlugins, []string{data.Name, cTags, string(rune(i))})
	}
	for i, author := range data.UserIds {
		yottadb.SetValE(yottadb.NOTTP, nil, string(rune(author)), cPlugins, []string{data.Name, cAuthors, string(rune(i))})
	}
	return nil
}

func PluginIdByName(name string) (int, error) { // TODO: Change name to exists()
	return 0, nil
}

func ListPlugins() ([]PluginData, error) {
	var currSub = ""
	var err error
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
		log.Println(currSub)
	}
	return nil, nil
}

func ListPluginsSearch(value string) ([]PluginData, error) {
	return nil, nil
}

func ListPluginsByUser(username string) ([]PluginData, error) {
	return nil, nil
}
