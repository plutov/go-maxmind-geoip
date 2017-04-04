#!/bin/bash

go get github.com/oschwald/geoip2-golang
go get -u github.com/jteeuwen/go-bindata/...
go-bindata -o geoip2-city.go geoip2-city.mmdb
go build -buildmode=plugin -o go-maxmind-geoip.so go-maxmind-geoip.go geoip2-city.go