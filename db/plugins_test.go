package main

import (
	"github.com/google/go-cmp/cmp"
	"testing"
)

var expected = PluginData{
	Name:             "Plugin0",
	ShortDescription: "None",
	Tags:             []string{"Tag0", "Tag1"},
	UserIds:          []int{0, 1},
}

func TestMain(m *testing.M) {
	_ = addPluginImpl(expected)
	_ = addPluginImpl(PluginData{
		Name:             "Plugin1",
		ShortDescription: "None",
		Tags:             []string{"Tag2", "Tag3"},
		UserIds:          []int{2},
	})
	m.Run()
}

func TestListPlugins_ByName(t *testing.T) {
	got, err := listPluginsImpl(ListRequest{
		Search: "Plugin0",
		UserId: -1,
	})
	if err != nil {
		t.Fatalf("List returned an error: %v\n", err)
	}
	if !cmp.Equal(expected, got) {
		t.Fatalf("Data is not equal! got: %v, expected: %v\n", got, expected)
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
	if !cmp.Equal(expected, got) {
		t.Fatalf("Data is not equal! got: %v, expected: %v\n", got, expected)
	}
}
