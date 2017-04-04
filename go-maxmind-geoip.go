package main

import (
	"github.com/oschwald/geoip2-golang"
	"sync"
)

// Client struct
type Client struct {
	db    *geoip2.Reader
	store map[string]*geoip2.City
	m     *sync.RWMutex
}

// NewClient func
func NewClient() (*Client, error) {
	dbBytes, err := Asset("geoip2-city.mmdb")
	if err != nil {
		return nil, err
	}

	db, err := geoip2.FromBytes(dbBytes)
	if err != nil {
		return nil, err
	}

	g := new(Client)
	g.db = db
	g.m = new(sync.RWMutex)
	g.store = make(map[string]*geoip2.City)

	return g, nil
}
