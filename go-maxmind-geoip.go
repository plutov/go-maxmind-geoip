package main

import (
	"github.com/oschwald/geoip2-golang"
	"net"
	"net/http"
)

var db *geoip2.Reader

func initDB() error {
	if db != nil {
		return nil
	}

	dbBytes, err := Asset("geoip2-city.mmdb")
	if err != nil {
		return err
	}

	var dbErr error
	db, dbErr = geoip2.FromBytes(dbBytes)
	return dbErr
}

func GetCity(r *http.Request) (string, error) {
	err := initDB()
	if err != nil {
		return "", err
	}

	ipStr := r.Header.Get("X-Real-IP")
	if len(ipStr) == 0 {
		ipStr = r.Header.Get("X-FORWARDED-FOR")
	}
	if len(ipStr) == 0 {
		ipStr, _, _ = net.SplitHostPort(r.RemoteAddr)
	}

	ip := net.ParseIP(ipStr)
	if ip == nil {
		return "", nil
	}

	record, err := db.City(ip)
	if err != nil {
		return "", err
	}

	return record.City.Names["en"], nil
}
