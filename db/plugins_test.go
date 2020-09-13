package main

import (
	"github.com/google/go-cmp/cmp"
	"lang.yottadb.com/go/yottadb"
	"testing"
)

var expected = PluginData{
	Name:             "Plugin0",
	ShortDescription: "None",
	Tags:             []string{"Tag0", "Tag1"},
	UserIds:          []int{0, 1},
}

func TestMain(m *testing.M) {
	// Delete db
	_ = yottadb.DeleteE(yottadb.NOTTP, nil, yottadb.YDB_DEL_TREE, "^plugins", []string{})

	// Add Data
	_ = addPluginImpl(expected)
	_ = addPluginImpl(PluginData{
		Name:             "Plugin1",
		ShortDescription: "None",
		Tags:             []string{"Tag2", "Tag3"},
		UserIds:          []int{2},
	})

	// Run
	m.Run()
}

func TestListPlugins_ByName(t *testing.T) {
	got, err := listPluginsImpl(ListRequest{
		Search: "plugin0",
		UserId: -1,
	})
	if err != nil {
		t.Fatalf("List returned an error: %v\n", err)
	}
	if len(got) != 1 {
		t.Fatalf("List size mismatch: want %v, got %v\n", 1, len(got))
	}
	if diff := cmp.Diff(expected, *(got[0])); diff != "" {
		t.Fatalf("listPluginsInfo() mismatch (-want +got):\n%s\n", diff)
	}
}

func TestListPlugins_ByUserId(t *testing.T) {
	got, err := listPluginsImpl(ListRequest{
		Search: "Plugin0",
		UserId: 0,
	})
	if err != nil {
		t.Fatalf("List returned an error: %v\n", err)
	}
	if len(got) != 1 {
		t.Fatalf("List size mismatch: want %v, got %v\n", 1, len(got))
	}
	if diff := cmp.Diff(expected, *(got[0])); diff != "" {
		t.Fatalf("listPluginsInfo() mismatch (-want +got):\n%s\n", diff)
	}
}
