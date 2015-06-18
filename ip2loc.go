package main

import (
  "flag"
  "os"
  "bufio"
  "fmt"
  "github.com/abh/geoip"
)

var db = flag.String("db", "/usr/share/GeoIP/GeoLiteCity.dat", "path to GeoLiteCity.dat")

func main() {
  flag.Parse()
  gi, err := geoip.Open(*db)
  if err != nil {
    fmt.Fprintf(os.Stderr, "%v\n", err)
    os.Exit(1)
  }

  scanner := bufio.NewScanner(os.Stdin)
  for scanner.Scan() {
    ip := scanner.Text()
    gr := gi.GetRecord(ip)
    country, city := "", ""
    if gr != nil {
      country, city = gr.CountryCode, gr.City
    }
    fmt.Printf("%v\t%v\t%v\n", ip, country, city)
  }
}
