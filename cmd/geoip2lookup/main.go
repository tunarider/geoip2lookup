package main

import (
	"fmt"
	"github.com/oschwald/geoip2-golang"
	"github.com/urfave/cli"
	"log"
	"net"
	"os"
)

func main() {
	app := cli.NewApp()
	app.Name = "geoip2lookup"
	app.Usage = "GeoIP2 lookup command line tool"
	app.Version = "0.1.0"
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "path, p",
			Value: "/usr/local/var/GeoIP/GeoLite2-City.mmdb",
			Usage: "Database file path",
		},
	}
	app.Action = func(ctx *cli.Context) (err error) {
		path := ctx.String("path")
		ipstr := ctx.Args().Get(0)
		db := readDatabase(path)
		defer db.Close()
		record := lookupIpv4(db, ipstr)
		printInfo(record)
		return err
	}
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

func readDatabase(path string) *geoip2.Reader {
	db, err := geoip2.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	return db
}

func lookupIpv4(db *geoip2.Reader, ipstr string) *geoip2.City {
	ip := net.ParseIP(ipstr)
	record, err := db.City(ip)
	if err != nil {
		log.Fatal(err)
	}
	return record
}

func printInfo(record *geoip2.City) {
	fmt.Printf("Country name: %v\n", record.Country.Names["en"])
	fmt.Printf("ISO country code: %v\n", record.Country.IsoCode)
	fmt.Printf("Timezone: %v\n", record.Location.TimeZone)
	fmt.Printf("Coordinates: %v, %v\n", record.Location.Latitude, record.Location.Longitude)
}
