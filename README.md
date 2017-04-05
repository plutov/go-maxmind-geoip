### For experimental usage only!

This [project](https://github.com/plutov/go-maxmind-geoip) contains an example with Go plugin which contains free GeoLite2 MaxMind's [database of ip addresses](http://dev.maxmind.com/geoip/geoip2/geolite2/).

It can find City by IP address.

It builds single `go-maxmind-geoip.so` plugin file with already included database with help of `go-bindata`.

### How to use in Go

Download `.so` plugin:
```
wget https://raw.githubusercontent.com/plutov/go-maxmind-geoip/master/go-maxmind-geoip.so
```

Use functions:
```
p, _ := plugin.Open("./go-maxmind-geoip.so")
init, _ := p.Lookup("InitDB")
init.(func() error)()
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
