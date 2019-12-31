# geoip2lookup

Command line tool for lookup GeoIP using GeoIP2 and GeoIPLite2 databases.

## Dependency

* [golang](https://golang.org)
* [dep](https://github.com/golang/dep)
* [MaxMind GeoLite2 Database](https://dev.maxmind.com/geoip/geoip2/geolite2/)

## Installation

```shell script
go get github.com/tunarider/geoip2lookup
cd $GOPATH/src/github.com/tunarider/geoip2lookup
make
make install
```

Be sure to check the `Makefile`.
The path is hardcoded.

## Usage

```shell script
NAME:
   geoip2lookup - GeoIP2 lookup command line tool

USAGE:
   geoip2lookup [global options] command [command options] [arguments...]

VERSION:
   0.2.0

COMMANDS:
   help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --path value, -p value  Database file path (default: "/usr/share/GeoIP/GeoLite2-City.mmdb")
   --help, -h              show help
   --version, -v           print the version
```
