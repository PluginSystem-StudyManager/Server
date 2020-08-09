package db

type PluginData struct {
	Name             string   `json:"name"`
	ShortDescription string   `json:"short_description"`
	Tags             []string `json:"tags"`
	UserIds          []int    `json:"user_ids"`
}
