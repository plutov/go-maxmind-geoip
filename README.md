### For experimental usage only!

This project contains an example with Go plugin which contains free GeoIP2 MaxMind's [database](http://dev.maxmind.com/geoip/geoip2/geolite2/).

It builds single `go-maxmind-geoip.so` with already included database with help of `go-bindata`.

### How to use in Go

Download `.so` plugin:
```
wget https://raw.githubusercontent.com/plutov/go-maxmind-geoip/master/go-maxmind-geoip.so
```

Use `GetCity` function:
```
p, _ := plugin.Open("./go-maxmind-geoip.so")
gc, _ := p.Lookup("GetCity")
city, _ := gc.(func(r *http.Request) (string, error))(r)
```

### Requirements

 - Linux or Darwin
 - Go 1.8

### Build

```
./build.sh
```

### Run tests

```
go test ./example/...
```