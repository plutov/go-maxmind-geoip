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

	gc, err := p.Lookup("GetCity")
	if err != nil {
		t.Fatalf("GetCity function not found. Details: %s", err.Error())
	}

	r, _ := http.NewRequest("GET", "/", nil)
	r.Header.Set("X-FORWARDED-FOR", "210.245.20.170")

	gcFunc, ok := gc.(func(r *http.Request) (string, error))
	if !ok {
		t.Fatal("Unable to cast func type")
	}

	city, err := gcFunc(r)
	if err != nil {
		t.Fatalf("Unable to find city. Details: %s", err.Error())
	}

	if city != "Hanoi" {
		t.Fatalf("Expected Hanoi, got %s", city)
	}
}
