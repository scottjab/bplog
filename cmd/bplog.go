package main

import (
	"flag"
	"github.com/scottjab/bplog"
	"log"
)

func main() {
	loadData := flag.String("load-data", "", "load data into db.")
	flag.Parse()

	b := bplog.NewBPLog()
	b.DBMigrate()

	if *loadData != "" {
		log.Printf("loading data from %s", *loadData)
		b.LoadCSVIntoDB(*loadData)
	}
	b.ServeHttp()
}

func loadData(string) {

}
