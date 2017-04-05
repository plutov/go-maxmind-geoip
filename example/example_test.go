package example

import (
	"net/http"
	"plugin"
	"testing"
)

func TestPlugin(t *testing.T) {
	pluginFile := "./../go-maxmind-geoip.so"
	p, err := plugin.Open(pluginFile)
	if err != nil {
		t.Fatalf("%s plugin file not found. Details: %s", pluginFile, err.Error())
	}

	init, err := p.Lookup("InitDB")
	if err != nil {
		t.Fatalf("InitDB function not found. Details: %s", err.Error())
	}

	initFunc, ok := init.(func() error)
	if !ok {
		t.Fatal("Unable to cast func type")
	}
	initFunc()

	gc, err := p.Lookup("GetCity")
	if err != nil {
		t.Fatalf("GetCity function not found. Details: %s", err.Error())
	}

	gcFunc, ok := gc.(func(r *http.Request) (string, error))
	if !ok {
		t.Fatal("Unable to cast func type")
	}

	r, _ := http.NewRequest("GET", "/", nil)
	r.Header.Set("X-FORWARDED-FOR", "210.245.20.170")
	city, err := gcFunc(r)
	if err != nil {
		t.Fatalf("Unable to find city. Details: %s", err.Error())
	}

	if city != "Hanoi" {
		t.Fatalf("Expected Hanoi, got %s", city)
	}
}
