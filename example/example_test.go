package example

import (
	"plugin"
	"testing"
)

func TestPlugin(t *testing.T) {
	pluginFile := "./../go-maxmind-geoip.so"
	p, err := plugin.Open(pluginFile)
	if err != nil {
		t.Fatalf("%s plugin file not found. Details: %s", pluginFile, err.Error())
	}

	_, err = p.Lookup("NewClient")
	if err != nil {
		t.Fatalf("NewClient function not found. Details: %s", err.Error())
	}
}
