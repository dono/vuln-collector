package main

import (
	"flag"
	"log"

	collector "github.com/dono/vuln-collector"
)

func main() {
	var (
		out = flag.String("out", "./vulns.sqlite3", "dump sqlite path")
	)
	flag.Parse()

	err := collector.DumpDB(*out)
	if err != nil {
		log.Fatal(err)
	}
}
