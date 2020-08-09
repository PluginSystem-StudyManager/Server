package db

import (
	"errors"
	"log"
)

type PluginData struct {
	Name             string   `json:"name"`
	ShortDescription string   `json:"short_description"`
	Tags             []string `json:"tags"`
	UserIds          []int    `json:"user_ids"`
}

func AddPlugin(data PluginData) error {
	res, err := insert("INSERT INTO plugins(Name, shortDescription) values(?, ?)", data.Name, data.ShortDescription)
	if err != nil {
		log.Fatal(err)
		return err
	}
	pluginId, err := res.LastInsertId()
	for _, tag := range data.Tags {
		res, err = insert("INSERT INTO plugins_tags(plugin, tag) values(?, ?)", pluginId, tag)
		if err != nil {
			log.Fatal(err)
			return err
		}
	}
	for _, userId := range data.UserIds {
		res, err = insert("INSERT INTO user_plugins(user, plugin) values(?, ?)", userId, pluginId)
		if err != nil {
			log.Fatal(err)
			return err
		}
	}
	return nil
}

func PluginIdByName(name string) (int, error) {
	rows, err := db.Query(`SELECT id FROM plugins WHERE name=?`, name)
	if err != nil {
		log.Fatal(err)
		return -1, err
	}
	defer rows.Close()
	for rows.Next() {
		var id int
		err = rows.Scan(&id)
		if err != nil {
			log.Fatal(err)
			return -1, err
		}
		return id, nil
	}
	return -1, errors.New("no valid token found: " + name)
}

func ListPlugins() ([]PluginData, error) {
	return listPluginsWhere("")
}

func ListPluginsSearch(value string) ([]PluginData, error) {
	return listPluginsWhere("WHERE UPPER(name) LIKE UPPER('%" + value + "%') OR UPPER(shortDescription) LIKE UPPER('%" + value + "%')")
}

func listPluginsWhere(where string) ([]PluginData, error) {
	rows, err := db.Query(`
	SELECT name, shortDescription
	FROM plugins 
	` + where + ";")
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	defer rows.Close()
	var plugins []PluginData
	for rows.Next() {
		var plugin PluginData
		err = rows.Scan(&plugin.Name, &plugin.ShortDescription)
		if err != nil {
			log.Fatal(err)
			return nil, err
		}
		plugins = append(plugins, plugin)
	}
	return plugins, nil
}
