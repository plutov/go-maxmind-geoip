### For experimental usage only!

This project contains an example with Go plugin which contains free GeoIP2 MaxMind's [database](http://dev.maxmind.com/geoip/geoip2/geolite2/).

### How to use in Go

```
wget https://raw.githubusercontent.com/plutov/go-maxmind-geoip/master/go-maxmind-geoip.so
```

```
p, _ := plugin.Open("./go-maxmind-geoip.so")  
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