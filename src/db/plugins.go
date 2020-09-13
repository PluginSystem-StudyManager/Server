package db

//go:generate schema-generate -o plugins.schema.go -p db ../../schemas/plugins/pluginData.schema.json ../../schemas/plugins/list.schema.json ../../schemas/plugins/add.schema.json

func Dummy() {
	// Needed that go:generate is executed
}
